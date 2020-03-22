# grpc-proxy

grpc反向代理

## todo
- 下游连接池
- 下游设置参数
- 基准测试
- 负载均衡

## 使用举例

```
// Binary server is an example server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/e421083458/grpc-proxy"
	_ "github.com/e421083458/grpc-proxy/encoding/gzip" // Install the gzip compressor
)

var port = flag.Int("port", 50051, "the port to serve on")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc.ProxyRealServerList = []string{"127.0.0.1:50055"}
	fmt.Printf("server listening at %v\n", lis.Addr())
	s := grpc.NewServer()
	s.Serve(lis)
}

```
