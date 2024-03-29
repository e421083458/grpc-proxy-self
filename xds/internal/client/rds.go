/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/e421083458/grpc-proxy/grpclog"

	xdspb "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

// newRDSRequest generates an RDS request proto for the provided routeName, to
// be sent out on the wire.
func (v2c *v2Client) newRDSRequest(routeName []string) *xdspb.DiscoveryRequest {
	return &xdspb.DiscoveryRequest{
		Node:          v2c.nodeProto,
		TypeUrl:       routeURL,
		ResourceNames: routeName,
	}
}

// sendRDS sends an RDS request for provided routeName on the provided stream.
func (v2c *v2Client) sendRDS(stream adsStream, routeName []string) bool {
	if err := stream.Send(v2c.newRDSRequest(routeName)); err != nil {
		grpclog.Infof("xds: RDS request for resource %v failed: %v", routeName, err)
		return false
	}
	return true
}

// handleRDSResponse processes an RDS response received from the xDS server. On
// receipt of a good response, it caches validated resources and also invokes
// the registered watcher callback.
func (v2c *v2Client) handleRDSResponse(resp *xdspb.DiscoveryResponse) error {
	v2c.mu.Lock()
	defer v2c.mu.Unlock()

	if v2c.watchMap[ldsResource] == nil {
		return fmt.Errorf("xds: unexpected RDS response when no LDS watcher is registered: %+v", resp)
	}
	target := v2c.watchMap[ldsResource].target[0]

	wi := v2c.watchMap[rdsResource]
	if wi == nil {
		return fmt.Errorf("xds: no RDS watcher found when handling RDS response: %+v", resp)
	}

	returnCluster := ""
	localCache := make(map[string]string)
	for _, r := range resp.GetResources() {
		var resource ptypes.DynamicAny
		if err := ptypes.UnmarshalAny(r, &resource); err != nil {
			return fmt.Errorf("xds: failed to unmarshal resource in RDS response: %v", err)
		}
		rc, ok := resource.Message.(*xdspb.RouteConfiguration)
		if !ok {
			return fmt.Errorf("xds: unexpected resource type: %T in RDS response", resource.Message)
		}
		cluster := getClusterFromRouteConfiguration(rc, target)
		if cluster == "" {
			return fmt.Errorf("xds: received invalid RouteConfiguration in RDS response: %+v", rc)
		}

		// If we get here, it means that this resource was a good one.
		localCache[rc.GetName()] = cluster
		if rc.GetName() == wi.target[0] {
			returnCluster = cluster
		}
	}

	// Update the cache in the v2Client only after we have confirmed that all
	// resources in the received response were good.
	for k, v := range localCache {
		// TODO: Need to handle deletion of entries from the cache based on LDS
		// watch calls. Not handling it does not affect correctness, but leads
		// to unnecessary memory consumption.
		v2c.rdsCache[k] = v
	}

	if returnCluster != "" {
		// We stop the expiry timer and invoke the callback only when we have
		// received the resource that we are watching for. Since RDS is an
		// incremental protocol, the fact that we did not receive the resource
		// that we are watching for in this response does not mean that the
		// server does not know about it.
		wi.expiryTimer.Stop()
		wi.callback.(rdsCallback)(rdsUpdate{clusterName: returnCluster}, nil)
	}
	return nil
}

// getClusterFromRouteConfiguration checks if the provided RouteConfiguration
// meets the expected criteria. If so, it returns a non-empty clusterName.
//
// A RouteConfiguration resource is considered valid when only if it contains a
// VirtualHost whose domain field matches the server name from the URI passed
// to the gRPC channel, and it contains a clusterName.
//
// The RouteConfiguration includes a list of VirtualHosts, which may have zero
// or more elements. We are interested in the element whose domains field
// matches the server name specified in the "xds:" URI (with port, if any,
// stripped off). The only field in the VirtualHost proto that the we are
// interested in is the list of routes. We only look at the last route in the
// list (the default route), whose match field must be empty and whose route
// field must be set.  Inside that route message, the cluster field will
// contain the clusterName we are looking for.
func getClusterFromRouteConfiguration(rc *xdspb.RouteConfiguration, target string) string {
	host, _, err := net.SplitHostPort(target)
	if err != nil {
		return ""
	}
	for _, vh := range rc.GetVirtualHosts() {
		for _, domain := range vh.GetDomains() {
			// TODO: Add support for wildcard matching here?
			if domain != host || len(vh.GetRoutes()) == 0 {
				continue
			}
			dr := vh.Routes[len(vh.Routes)-1]
			if dr.GetMatch() == nil && dr.GetRoute() != nil {
				return dr.GetRoute().GetCluster()
			}
		}
	}
	return ""
}
