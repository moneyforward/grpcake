// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: foo/example.proto

package foo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExampleService_UnaryExample_FullMethodName               = "/foo.ExampleService/UnaryExample"
	ExampleService_ClientStreamExample_FullMethodName        = "/foo.ExampleService/ClientStreamExample"
	ExampleService_ServerStreamExample_FullMethodName        = "/foo.ExampleService/ServerStreamExample"
	ExampleService_BidirectionalStreamExample_FullMethodName = "/foo.ExampleService/BidirectionalStreamExample"
	ExampleService_AdvancedExample_FullMethodName            = "/foo.ExampleService/AdvancedExample"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// Unary RPC
	UnaryExample(ctx context.Context, in *BasicTypes, opts ...grpc.CallOption) (*BasicTypes, error)
	// Client streaming RPC
	ClientStreamExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService_ClientStreamExampleClient, error)
	// Server streaming RPC
	ServerStreamExample(ctx context.Context, in *BasicTypes, opts ...grpc.CallOption) (ExampleService_ServerStreamExampleClient, error)
	// Bi-directional streaming RPC
	BidirectionalStreamExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService_BidirectionalStreamExampleClient, error)
	// RPC with advanced types
	AdvancedExample(ctx context.Context, in *AdvancedTypes, opts ...grpc.CallOption) (*AdvancedTypes, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) UnaryExample(ctx context.Context, in *BasicTypes, opts ...grpc.CallOption) (*BasicTypes, error) {
	out := new(BasicTypes)
	err := c.cc.Invoke(ctx, ExampleService_UnaryExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ClientStreamExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService_ClientStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_ClientStreamExample_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceClientStreamExampleClient{stream}
	return x, nil
}

type ExampleService_ClientStreamExampleClient interface {
	Send(*BasicTypes) error
	CloseAndRecv() (*BasicTypes, error)
	grpc.ClientStream
}

type exampleServiceClientStreamExampleClient struct {
	grpc.ClientStream
}

func (x *exampleServiceClientStreamExampleClient) Send(m *BasicTypes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceClientStreamExampleClient) CloseAndRecv() (*BasicTypes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BasicTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) ServerStreamExample(ctx context.Context, in *BasicTypes, opts ...grpc.CallOption) (ExampleService_ServerStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_ServerStreamExample_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceServerStreamExampleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_ServerStreamExampleClient interface {
	Recv() (*BasicTypes, error)
	grpc.ClientStream
}

type exampleServiceServerStreamExampleClient struct {
	grpc.ClientStream
}

func (x *exampleServiceServerStreamExampleClient) Recv() (*BasicTypes, error) {
	m := new(BasicTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) BidirectionalStreamExample(ctx context.Context, opts ...grpc.CallOption) (ExampleService_BidirectionalStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[2], ExampleService_BidirectionalStreamExample_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceBidirectionalStreamExampleClient{stream}
	return x, nil
}

type ExampleService_BidirectionalStreamExampleClient interface {
	Send(*BasicTypes) error
	Recv() (*BasicTypes, error)
	grpc.ClientStream
}

type exampleServiceBidirectionalStreamExampleClient struct {
	grpc.ClientStream
}

func (x *exampleServiceBidirectionalStreamExampleClient) Send(m *BasicTypes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceBidirectionalStreamExampleClient) Recv() (*BasicTypes, error) {
	m := new(BasicTypes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) AdvancedExample(ctx context.Context, in *AdvancedTypes, opts ...grpc.CallOption) (*AdvancedTypes, error) {
	out := new(AdvancedTypes)
	err := c.cc.Invoke(ctx, ExampleService_AdvancedExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	// Unary RPC
	UnaryExample(context.Context, *BasicTypes) (*BasicTypes, error)
	// Client streaming RPC
	ClientStreamExample(ExampleService_ClientStreamExampleServer) error
	// Server streaming RPC
	ServerStreamExample(*BasicTypes, ExampleService_ServerStreamExampleServer) error
	// Bi-directional streaming RPC
	BidirectionalStreamExample(ExampleService_BidirectionalStreamExampleServer) error
	// RPC with advanced types
	AdvancedExample(context.Context, *AdvancedTypes) (*AdvancedTypes, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) UnaryExample(context.Context, *BasicTypes) (*BasicTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryExample not implemented")
}
func (UnimplementedExampleServiceServer) ClientStreamExample(ExampleService_ClientStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamExample not implemented")
}
func (UnimplementedExampleServiceServer) ServerStreamExample(*BasicTypes, ExampleService_ServerStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamExample not implemented")
}
func (UnimplementedExampleServiceServer) BidirectionalStreamExample(ExampleService_BidirectionalStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamExample not implemented")
}
func (UnimplementedExampleServiceServer) AdvancedExample(context.Context, *AdvancedTypes) (*AdvancedTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvancedExample not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_UnaryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicTypes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).UnaryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_UnaryExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).UnaryExample(ctx, req.(*BasicTypes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ClientStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).ClientStreamExample(&exampleServiceClientStreamExampleServer{stream})
}

type ExampleService_ClientStreamExampleServer interface {
	SendAndClose(*BasicTypes) error
	Recv() (*BasicTypes, error)
	grpc.ServerStream
}

type exampleServiceClientStreamExampleServer struct {
	grpc.ServerStream
}

func (x *exampleServiceClientStreamExampleServer) SendAndClose(m *BasicTypes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceClientStreamExampleServer) Recv() (*BasicTypes, error) {
	m := new(BasicTypes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExampleService_ServerStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BasicTypes)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).ServerStreamExample(m, &exampleServiceServerStreamExampleServer{stream})
}

type ExampleService_ServerStreamExampleServer interface {
	Send(*BasicTypes) error
	grpc.ServerStream
}

type exampleServiceServerStreamExampleServer struct {
	grpc.ServerStream
}

func (x *exampleServiceServerStreamExampleServer) Send(m *BasicTypes) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService_BidirectionalStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).BidirectionalStreamExample(&exampleServiceBidirectionalStreamExampleServer{stream})
}

type ExampleService_BidirectionalStreamExampleServer interface {
	Send(*BasicTypes) error
	Recv() (*BasicTypes, error)
	grpc.ServerStream
}

type exampleServiceBidirectionalStreamExampleServer struct {
	grpc.ServerStream
}

func (x *exampleServiceBidirectionalStreamExampleServer) Send(m *BasicTypes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceBidirectionalStreamExampleServer) Recv() (*BasicTypes, error) {
	m := new(BasicTypes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExampleService_AdvancedExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvancedTypes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).AdvancedExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_AdvancedExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).AdvancedExample(ctx, req.(*AdvancedTypes))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foo.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryExample",
			Handler:    _ExampleService_UnaryExample_Handler,
		},
		{
			MethodName: "AdvancedExample",
			Handler:    _ExampleService_AdvancedExample_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamExample",
			Handler:       _ExampleService_ClientStreamExample_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamExample",
			Handler:       _ExampleService_ServerStreamExample_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamExample",
			Handler:       _ExampleService_BidirectionalStreamExample_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "foo/example.proto",
}
