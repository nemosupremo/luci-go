// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/grpc/discovery/internal/testservices/helloworld.proto

package testservices

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_c619609dd02e7cb1, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_c619609dd02e7cb1, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MultiplyRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyRequest) Reset()         { *m = MultiplyRequest{} }
func (m *MultiplyRequest) String() string { return proto.CompactTextString(m) }
func (*MultiplyRequest) ProtoMessage()    {}
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_c619609dd02e7cb1, []int{2}
}
func (m *MultiplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyRequest.Unmarshal(m, b)
}
func (m *MultiplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyRequest.Marshal(b, m, deterministic)
}
func (dst *MultiplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyRequest.Merge(dst, src)
}
func (m *MultiplyRequest) XXX_Size() int {
	return xxx_messageInfo_MultiplyRequest.Size(m)
}
func (m *MultiplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyRequest proto.InternalMessageInfo

func (m *MultiplyRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *MultiplyRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type MultiplyResponse struct {
	Z                    int32    `protobuf:"varint,1,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplyResponse) Reset()         { *m = MultiplyResponse{} }
func (m *MultiplyResponse) String() string { return proto.CompactTextString(m) }
func (*MultiplyResponse) ProtoMessage()    {}
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_c619609dd02e7cb1, []int{3}
}
func (m *MultiplyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplyResponse.Unmarshal(m, b)
}
func (m *MultiplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplyResponse.Marshal(b, m, deterministic)
}
func (dst *MultiplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplyResponse.Merge(dst, src)
}
func (m *MultiplyResponse) XXX_Size() int {
	return xxx_messageInfo_MultiplyResponse.Size(m)
}
func (m *MultiplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplyResponse proto.InternalMessageInfo

func (m *MultiplyResponse) GetZ() int32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "testservices.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "testservices.HelloReply")
	proto.RegisterType((*MultiplyRequest)(nil), "testservices.MultiplyRequest")
	proto.RegisterType((*MultiplyResponse)(nil), "testservices.MultiplyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
type greeterPRPCClient struct {
	client *prpc.Client
}

func NewGreeterPRPCClient(client *prpc.Client) GreeterClient {
	return &greeterPRPCClient{client}
}

func (c *greeterPRPCClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.client.Call(ctx, "testservices.Greeter", "SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/testservices.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s prpc.Registrar, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testservices.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testservices.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/discovery/internal/testservices/helloworld.proto",
}

// CalcClient is the client API for Calc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcClient interface {
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
}
type calcPRPCClient struct {
	client *prpc.Client
}

func NewCalcPRPCClient(client *prpc.Client) CalcClient {
	return &calcPRPCClient{client}
}

func (c *calcPRPCClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.client.Call(ctx, "testservices.Calc", "Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type calcClient struct {
	cc *grpc.ClientConn
}

func NewCalcClient(cc *grpc.ClientConn) CalcClient {
	return &calcClient{cc}
}

func (c *calcClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, "/testservices.Calc/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServer is the server API for Calc service.
type CalcServer interface {
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
}

func RegisterCalcServer(s prpc.Registrar, srv CalcServer) {
	s.RegisterService(&_Calc_serviceDesc, srv)
}

func _Calc_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testservices.Calc/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testservices.Calc",
	HandlerType: (*CalcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _Calc_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/discovery/internal/testservices/helloworld.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/grpc/discovery/internal/testservices/helloworld.proto", fileDescriptor_helloworld_c619609dd02e7cb1)
}

var fileDescriptor_helloworld_c619609dd02e7cb1 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x57, 0x99, 0x6e, 0x1e, 0x0a, 0x4a, 0xae, 0x4a, 0x41, 0x19, 0xb9, 0x10, 0x6f, 0x6c,
	0x60, 0xbe, 0x81, 0x5e, 0x28, 0xca, 0x6e, 0xba, 0x27, 0xa8, 0xd9, 0xa1, 0x0b, 0xa4, 0x4d, 0x3c,
	0x49, 0xe7, 0xb2, 0xa7, 0x97, 0x76, 0x06, 0xab, 0xe8, 0x5d, 0x7e, 0xfe, 0x8f, 0x9c, 0xef, 0x1c,
	0x78, 0xa9, 0x4d, 0x21, 0xb7, 0x64, 0x1a, 0xd5, 0x35, 0x85, 0xa1, 0x5a, 0xe8, 0x4e, 0x2a, 0x51,
	0x93, 0x95, 0x62, 0xa3, 0x9c, 0x34, 0x3b, 0xa4, 0x20, 0x54, 0xeb, 0x91, 0xda, 0x4a, 0x0b, 0x8f,
	0xce, 0x3b, 0xa4, 0x9d, 0x92, 0xe8, 0xc4, 0x16, 0xb5, 0x36, 0x1f, 0x86, 0xf4, 0xa6, 0xb0, 0x64,
	0xbc, 0x61, 0xe9, 0xb8, 0xe6, 0x1c, 0xd2, 0xe7, 0x9e, 0x28, 0xf1, 0xbd, 0x43, 0xe7, 0x19, 0x83,
	0x69, 0x5b, 0x35, 0x98, 0x25, 0x8b, 0xe4, 0xf6, 0xbc, 0x1c, 0xde, 0xfc, 0x06, 0xe0, 0x8b, 0xb1,
	0x3a, 0xb0, 0x0c, 0x66, 0x0d, 0x3a, 0x57, 0xd5, 0x11, 0x8a, 0x91, 0xdf, 0xc1, 0xc5, 0xaa, 0xd3,
	0x5e, 0x59, 0x1d, 0xe2, 0x77, 0x29, 0x24, 0xfb, 0x01, 0x3b, 0x2d, 0x93, 0x7d, 0x9f, 0x42, 0x76,
	0x72, 0x4c, 0x81, 0x2f, 0xe0, 0xf2, 0x1b, 0x77, 0xd6, 0xb4, 0x0e, 0x7b, 0xe2, 0x10, 0xf9, 0xc3,
	0x72, 0x05, 0xb3, 0x27, 0x42, 0xf4, 0x48, 0xec, 0x01, 0xe6, 0xeb, 0x2a, 0x0c, 0x1a, 0x2c, 0x2f,
	0xc6, 0x2b, 0x14, 0x63, 0xff, 0x3c, 0xfb, 0xb3, 0xb3, 0x3a, 0xf0, 0xc9, 0x72, 0x0d, 0xd3, 0xc7,
	0x4a, 0x4b, 0xf6, 0x0a, 0xf3, 0x38, 0x98, 0x5d, 0xfd, 0xe4, 0x7f, 0xf9, 0xe7, 0xd7, 0xff, 0xd5,
	0x47, 0x5f, 0x3e, 0x79, 0x3b, 0x1b, 0xae, 0x7a, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x1a,
	0xdd, 0x06, 0xa3, 0x01, 0x00, 0x00,
}
