// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	QueryRequest
	QueryResponse
	ExecRequest
	ExecResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// The query statement
type QueryRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto1.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The query response
type QueryResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto1.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The exec statement
type ExecRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExecRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The exec response
type ExecResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto1.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExecResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*QueryRequest)(nil), "proto.QueryRequest")
	proto1.RegisterType((*QueryResponse)(nil), "proto.QueryResponse")
	proto1.RegisterType((*ExecRequest)(nil), "proto.ExecRequest")
	proto1.RegisterType((*ExecResponse)(nil), "proto.ExecResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DBProvider service

type DBProviderClient interface {
	// Query executes a statement that reads data.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type dBProviderClient struct {
	cc *grpc.ClientConn
}

func NewDBProviderClient(cc *grpc.ClientConn) DBProviderClient {
	return &dBProviderClient{cc}
}

func (c *dBProviderClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBProviderClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DBProvider service

type DBProviderServer interface {
	// Query executes a statement that reads data.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

func RegisterDBProviderServer(s *grpc.Server, srv DBProviderServer) {
	s.RegisterService(&_DBProvider_serviceDesc, srv)
}

func _DBProvider_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBProvider_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DBProvider",
	HandlerType: (*DBProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DBProvider_Query_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _DBProvider_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x4a, 0x5c,
	0x3c, 0x81, 0xa5, 0xa9, 0x45, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c,
	0x2c, 0xc5, 0x25, 0xb9, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x26,
	0x17, 0x2f, 0x54, 0x4d, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a,
	0x71, 0x71, 0x62, 0x7a, 0x2a, 0x54, 0x1d, 0x8c, 0xab, 0xa4, 0xc8, 0xc5, 0xed, 0x5a, 0x91, 0x9a,
	0x8c, 0xcf, 0x34, 0x0d, 0x2e, 0x1e, 0x88, 0x12, 0x42, 0x86, 0x19, 0x95, 0x72, 0x71, 0xb9, 0x38,
	0x05, 0x14, 0xe5, 0x97, 0x65, 0xa6, 0xa4, 0x16, 0x09, 0x99, 0x70, 0xb1, 0x82, 0x5d, 0x21, 0x24,
	0x0c, 0xf1, 0x81, 0x1e, 0xb2, 0xbb, 0xa5, 0x44, 0x50, 0x05, 0x21, 0x66, 0x2b, 0x31, 0x08, 0x19,
	0x72, 0xb1, 0x80, 0x6c, 0x13, 0x12, 0x82, 0xca, 0x23, 0xb9, 0x4e, 0x4a, 0x18, 0x45, 0x0c, 0xa6,
	0x25, 0x89, 0x0d, 0x2c, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x53, 0xff, 0x7a, 0x56, 0x31,
	0x01, 0x00, 0x00,
}
