// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_testing/test.proto

package grpc_testing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "github.com/e421083458/grpc-proxy"
	codes "github.com/e421083458/grpc-proxy/codes"
	status "github.com/e421083458/grpc-proxy/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The type of payload that should be returned.
type PayloadType int32

const (
	// Compressable text format.
	PayloadType_COMPRESSABLE PayloadType = 0
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "COMPRESSABLE",
	1: "UNCOMPRESSABLE",
	2: "RANDOM",
}

var PayloadType_value = map[string]int32{
	"COMPRESSABLE":   0,
	"UNCOMPRESSABLE": 1,
	"RANDOM":         2,
}

func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}

func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// A block of data, to simply increase gRPC message size.
type Payload struct {
	// The type of data in body.
	Type PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{1}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetType() PayloadType {
	if m != nil {
		return m.Type
	}
	return PayloadType_COMPRESSABLE
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// Unary request.
type SimpleRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize int32 `protobuf:"varint,2,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Whether SimpleResponse should include username.
	FillUsername bool `protobuf:"varint,4,opt,name=fill_username,json=fillUsername,proto3" json:"fill_username,omitempty"`
	// Whether SimpleResponse should include OAuth scope.
	FillOauthScope       bool     `protobuf:"varint,5,opt,name=fill_oauth_scope,json=fillOauthScope,proto3" json:"fill_oauth_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{2}
}

func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

func (m *SimpleRequest) GetResponseType() PayloadType {
	if m != nil {
		return m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *SimpleRequest) GetResponseSize() int32 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *SimpleRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleRequest) GetFillUsername() bool {
	if m != nil {
		return m.FillUsername
	}
	return false
}

func (m *SimpleRequest) GetFillOauthScope() bool {
	if m != nil {
		return m.FillOauthScope
	}
	return false
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	// Payload to increase message size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// The user the request came from, for verifying authentication was
	// successful when the client expected it.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// OAuth scope.
	OauthScope           string   `protobuf:"bytes,3,opt,name=oauth_scope,json=oauthScope,proto3" json:"oauth_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{3}
}

func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SimpleResponse) GetOauthScope() string {
	if m != nil {
		return m.OauthScope
	}
	return ""
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	// Optional input payload sent along with the request.
	Payload              *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingInputCallRequest) Reset()         { *m = StreamingInputCallRequest{} }
func (m *StreamingInputCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingInputCallRequest) ProtoMessage()    {}
func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{4}
}

func (m *StreamingInputCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingInputCallRequest.Unmarshal(m, b)
}
func (m *StreamingInputCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingInputCallRequest.Marshal(b, m, deterministic)
}
func (m *StreamingInputCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingInputCallRequest.Merge(m, src)
}
func (m *StreamingInputCallRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingInputCallRequest.Size(m)
}
func (m *StreamingInputCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingInputCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingInputCallRequest proto.InternalMessageInfo

func (m *StreamingInputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize int32    `protobuf:"varint,1,opt,name=aggregated_payload_size,json=aggregatedPayloadSize,proto3" json:"aggregated_payload_size,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *StreamingInputCallResponse) Reset()         { *m = StreamingInputCallResponse{} }
func (m *StreamingInputCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingInputCallResponse) ProtoMessage()    {}
func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{5}
}

func (m *StreamingInputCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingInputCallResponse.Unmarshal(m, b)
}
func (m *StreamingInputCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingInputCallResponse.Marshal(b, m, deterministic)
}
func (m *StreamingInputCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingInputCallResponse.Merge(m, src)
}
func (m *StreamingInputCallResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingInputCallResponse.Size(m)
}
func (m *StreamingInputCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingInputCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingInputCallResponse proto.InternalMessageInfo

func (m *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if m != nil {
		return m.AggregatedPayloadSize
	}
	return 0
}

// Configuration for a particular response.
type ResponseParameters struct {
	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs           int32    `protobuf:"varint,2,opt,name=interval_us,json=intervalUs,proto3" json:"interval_us,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseParameters) Reset()         { *m = ResponseParameters{} }
func (m *ResponseParameters) String() string { return proto.CompactTextString(m) }
func (*ResponseParameters) ProtoMessage()    {}
func (*ResponseParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{6}
}

func (m *ResponseParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseParameters.Unmarshal(m, b)
}
func (m *ResponseParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseParameters.Marshal(b, m, deterministic)
}
func (m *ResponseParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseParameters.Merge(m, src)
}
func (m *ResponseParameters) XXX_Size() int {
	return xxx_messageInfo_ResponseParameters.Size(m)
}
func (m *ResponseParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseParameters proto.InternalMessageInfo

func (m *ResponseParameters) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ResponseParameters) GetIntervalUs() int32 {
	if m != nil {
		return m.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, the payload from each response in the stream
	// might be of different types. This is to simulate a mixed type of payload
	// stream.
	ResponseType PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters,json=responseParameters,proto3" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	Payload              *Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingOutputCallRequest) Reset()         { *m = StreamingOutputCallRequest{} }
func (m *StreamingOutputCallRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingOutputCallRequest) ProtoMessage()    {}
func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{7}
}

func (m *StreamingOutputCallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOutputCallRequest.Unmarshal(m, b)
}
func (m *StreamingOutputCallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOutputCallRequest.Marshal(b, m, deterministic)
}
func (m *StreamingOutputCallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOutputCallRequest.Merge(m, src)
}
func (m *StreamingOutputCallRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingOutputCallRequest.Size(m)
}
func (m *StreamingOutputCallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOutputCallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOutputCallRequest proto.InternalMessageInfo

func (m *StreamingOutputCallRequest) GetResponseType() PayloadType {
	if m != nil {
		return m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if m != nil {
		return m.ResponseParameters
	}
	return nil
}

func (m *StreamingOutputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	// Payload to increase response size.
	Payload              *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingOutputCallResponse) Reset()         { *m = StreamingOutputCallResponse{} }
func (m *StreamingOutputCallResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingOutputCallResponse) ProtoMessage()    {}
func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cda82041fed8bf, []int{8}
}

func (m *StreamingOutputCallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingOutputCallResponse.Unmarshal(m, b)
}
func (m *StreamingOutputCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingOutputCallResponse.Marshal(b, m, deterministic)
}
func (m *StreamingOutputCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingOutputCallResponse.Merge(m, src)
}
func (m *StreamingOutputCallResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingOutputCallResponse.Size(m)
}
func (m *StreamingOutputCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingOutputCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingOutputCallResponse proto.InternalMessageInfo

func (m *StreamingOutputCallResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("grpc.testing.PayloadType", PayloadType_name, PayloadType_value)
	proto.RegisterType((*Empty)(nil), "grpc.testing.Empty")
	proto.RegisterType((*Payload)(nil), "grpc.testing.Payload")
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterType((*StreamingInputCallRequest)(nil), "grpc.testing.StreamingInputCallRequest")
	proto.RegisterType((*StreamingInputCallResponse)(nil), "grpc.testing.StreamingInputCallResponse")
	proto.RegisterType((*ResponseParameters)(nil), "grpc.testing.ResponseParameters")
	proto.RegisterType((*StreamingOutputCallRequest)(nil), "grpc.testing.StreamingOutputCallRequest")
	proto.RegisterType((*StreamingOutputCallResponse)(nil), "grpc.testing.StreamingOutputCallResponse")
}

func init() { proto.RegisterFile("grpc_testing/test.proto", fileDescriptor_e1cda82041fed8bf) }

var fileDescriptor_e1cda82041fed8bf = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0xdb, 0xf4, 0x36, 0x49, 0xad, 0x68, 0xab, 0xaa, 0xae, 0x8b, 0x84, 0x65, 0x1e, 0x30,
	0x48, 0xa4, 0x28, 0x08, 0x1e, 0x41, 0xa5, 0x17, 0x51, 0x29, 0x4d, 0x82, 0x9d, 0x3c, 0x47, 0xdb,
	0x64, 0x6b, 0x2c, 0x39, 0xf6, 0xb2, 0x5e, 0x57, 0xa4, 0x0f, 0xfc, 0x18, 0x3f, 0xc3, 0x47, 0xf0,
	0x01, 0x68, 0xd7, 0x76, 0xe2, 0x24, 0xae, 0x48, 0x41, 0xf0, 0x14, 0x7b, 0xe6, 0xcc, 0x99, 0x73,
	0x3c, 0xb3, 0x1b, 0x38, 0xf0, 0x38, 0x1b, 0x0e, 0x04, 0x8d, 0x85, 0x1f, 0x7a, 0xc7, 0xf2, 0xb7,
	0xc1, 0x78, 0x24, 0x22, 0x5c, 0x93, 0x89, 0x46, 0x96, 0xb0, 0xb6, 0x60, 0xe3, 0x7c, 0xcc, 0xc4,
	0xc4, 0x6a, 0xc1, 0x56, 0x97, 0x4c, 0x82, 0x88, 0x8c, 0xf0, 0x4b, 0xa8, 0x88, 0x09, 0xa3, 0x3a,
	0x32, 0x91, 0xad, 0x35, 0x0f, 0x1b, 0xc5, 0x82, 0x46, 0x06, 0xea, 0x4d, 0x18, 0x75, 0x14, 0x0c,
	0x63, 0xa8, 0x5c, 0x47, 0xa3, 0x89, 0xbe, 0x66, 0x22, 0xbb, 0xe6, 0xa8, 0x67, 0xeb, 0x27, 0x82,
	0x5d, 0xd7, 0x1f, 0xb3, 0x80, 0x3a, 0xf4, 0x4b, 0x42, 0x63, 0x81, 0xdf, 0xc1, 0x2e, 0xa7, 0x31,
	0x8b, 0xc2, 0x98, 0x0e, 0x56, 0x63, 0xaf, 0xe5, 0x78, 0xf9, 0x86, 0x9f, 0x16, 0xea, 0x63, 0xff,
	0x8e, 0xaa, 0x76, 0x1b, 0x33, 0x90, 0xeb, 0xdf, 0x51, 0x7c, 0x0c, 0x5b, 0x2c, 0x65, 0xd0, 0xd7,
	0x4d, 0x64, 0x57, 0x9b, 0xfb, 0xa5, 0xf4, 0x4e, 0x8e, 0x92, 0xac, 0x37, 0x7e, 0x10, 0x0c, 0x92,
	0x98, 0xf2, 0x90, 0x8c, 0xa9, 0x5e, 0x31, 0x91, 0xbd, 0xed, 0xd4, 0x64, 0xb0, 0x9f, 0xc5, 0xb0,
	0x0d, 0x75, 0x05, 0x8a, 0x48, 0x22, 0x3e, 0x0f, 0xe2, 0x61, 0xc4, 0xa8, 0xbe, 0xa1, 0x70, 0x9a,
	0x8c, 0x77, 0x64, 0xd8, 0x95, 0x51, 0xeb, 0x1b, 0x68, 0xb9, 0xeb, 0x54, 0x55, 0x51, 0x11, 0x5a,
	0x49, 0x91, 0x01, 0xdb, 0x53, 0x31, 0xd2, 0xe2, 0x8e, 0x33, 0x7d, 0xc7, 0x4f, 0xa0, 0x5a, 0xd4,
	0xb0, 0xae, 0xd2, 0x10, 0xcd, 0xfa, 0xb7, 0xe0, 0xd0, 0x15, 0x9c, 0x92, 0xb1, 0x1f, 0x7a, 0x97,
	0x21, 0x4b, 0xc4, 0x29, 0x09, 0x82, 0x7c, 0x02, 0x0f, 0x95, 0x62, 0xf5, 0xc0, 0x28, 0x63, 0xcb,
	0x9c, 0xbd, 0x85, 0x03, 0xe2, 0x79, 0x9c, 0x7a, 0x44, 0xd0, 0xd1, 0x20, 0xab, 0x49, 0x47, 0x83,
	0xd4, 0x68, 0xf6, 0x67, 0xe9, 0x8c, 0x5a, 0xce, 0xc8, 0xba, 0x04, 0x9c, 0x73, 0x74, 0x09, 0x27,
	0x63, 0x2a, 0x28, 0x8f, 0xe5, 0x12, 0x15, 0x4a, 0xd5, 0xb3, 0xb4, 0xeb, 0x87, 0x82, 0xf2, 0x5b,
	0x22, 0x07, 0x94, 0x0d, 0x1c, 0xf2, 0x50, 0x3f, 0xb6, 0x7e, 0xa0, 0x82, 0xc2, 0x4e, 0x22, 0x16,
	0x0c, 0xff, 0xed, 0xca, 0x7d, 0x82, 0xbd, 0x69, 0x3d, 0x9b, 0x4a, 0xd5, 0xd7, 0xcc, 0x75, 0xbb,
	0xda, 0x34, 0xe7, 0x59, 0x96, 0x2d, 0x39, 0x98, 0x2f, 0xdb, 0x7c, 0xe8, 0x82, 0x5a, 0x6d, 0x38,
	0x2a, 0x75, 0xf8, 0x87, 0xeb, 0xf5, 0xe2, 0x3d, 0x54, 0x0b, 0x86, 0x71, 0x1d, 0x6a, 0xa7, 0x9d,
	0xab, 0xae, 0x73, 0xee, 0xba, 0x27, 0x1f, 0x5a, 0xe7, 0xf5, 0x47, 0x18, 0x83, 0xd6, 0x6f, 0xcf,
	0xc5, 0x10, 0x06, 0xd8, 0x74, 0x4e, 0xda, 0x67, 0x9d, 0xab, 0xfa, 0x5a, 0xf3, 0x7b, 0x05, 0xaa,
	0x3d, 0x1a, 0x0b, 0x97, 0xf2, 0x5b, 0x7f, 0x48, 0xf1, 0x1b, 0xd8, 0x51, 0x17, 0x88, 0x94, 0x85,
	0xf7, 0xe6, 0xbb, 0xab, 0x84, 0x51, 0x16, 0xc4, 0x17, 0xb0, 0xd3, 0x0f, 0x09, 0x4f, 0xcb, 0x8e,
	0xe6, 0x11, 0x73, 0x17, 0x87, 0xf1, 0xb8, 0x3c, 0x99, 0x7d, 0x80, 0x00, 0xf6, 0x4a, 0xbe, 0x0f,
	0xb6, 0x17, 0x8a, 0xee, 0x5d, 0x12, 0xe3, 0xf9, 0x0a, 0xc8, 0xb4, 0xd7, 0x2b, 0x84, 0x7d, 0xc0,
	0xcb, 0x27, 0x02, 0x3f, 0xbb, 0x87, 0x62, 0xf1, 0x04, 0x1a, 0xf6, 0xef, 0x81, 0x69, 0x2b, 0x5b,
	0xb6, 0xd2, 0x2e, 0x92, 0x20, 0x38, 0x4b, 0x58, 0x40, 0xbf, 0xfe, 0x33, 0x4f, 0x36, 0x52, 0xae,
	0xb4, 0x8f, 0x24, 0xb8, 0xf9, 0x0f, 0xad, 0xae, 0x37, 0xd5, 0x7f, 0xd0, 0xeb, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x07, 0xc7, 0x76, 0x69, 0x9e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/e421083458/grpc-proxy#ClientConn.NewStream.
type TestServiceClient interface {
	// One empty request followed by one empty response.
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/EmptyCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/grpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingOutputCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamingOutputCallClient interface {
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingOutputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingOutputCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/grpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingInputCallClient{stream}
	return x, nil
}

type TestService_StreamingInputCallClient interface {
	Send(*StreamingInputCallRequest) error
	CloseAndRecv() (*StreamingInputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingInputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingInputCallClient) Send(m *StreamingInputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*StreamingInputCallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingInputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[3], "/grpc.testing.TestService/HalfDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceHalfDuplexCallClient{stream}
	return x, nil
}

type TestService_HalfDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceHalfDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceHalfDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	// One empty request followed by one empty response.
	EmptyCall(context.Context, *Empty) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(TestService_StreamingInputCallServer) error
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(TestService_HalfDuplexCallServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) EmptyCall(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCall not implemented")
}
func (*UnimplementedTestServiceServer) UnaryCall(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (*UnimplementedTestServiceServer) StreamingOutputCall(req *StreamingOutputCallRequest, srv TestService_StreamingOutputCallServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingOutputCall not implemented")
}
func (*UnimplementedTestServiceServer) StreamingInputCall(srv TestService_StreamingInputCallServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingInputCall not implemented")
}
func (*UnimplementedTestServiceServer) FullDuplexCall(srv TestService_FullDuplexCallServer) error {
	return status.Errorf(codes.Unimplemented, "method FullDuplexCall not implemented")
}
func (*UnimplementedTestServiceServer) HalfDuplexCall(srv TestService_HalfDuplexCallServer) error {
	return status.Errorf(codes.Unimplemented, "method HalfDuplexCall not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/EmptyCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingOutputCallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer{stream})
}

type TestService_StreamingOutputCallServer interface {
	Send(*StreamingOutputCallResponse) error
	grpc.ServerStream
}

type testServiceStreamingOutputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingOutputCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer{stream})
}

type TestService_StreamingInputCallServer interface {
	SendAndClose(*StreamingInputCallResponse) error
	Recv() (*StreamingInputCallRequest, error)
	grpc.ServerStream
}

type testServiceStreamingInputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingInputCallServer) SendAndClose(m *StreamingInputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallServer) Recv() (*StreamingInputCallRequest, error) {
	m := new(StreamingInputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_HalfDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).HalfDuplexCall(&testServiceHalfDuplexCallServer{stream})
}

type TestService_HalfDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceHalfDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceHalfDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_testing/test.proto",
}
