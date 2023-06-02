// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: composite/example.proto

package composite

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
	CompositeService_UnaryExample_FullMethodName = "/composite.CompositeService/UnaryExample"
)

// CompositeServiceClient is the client API for CompositeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompositeServiceClient interface {
	UnaryExample(ctx context.Context, in *CompositeType, opts ...grpc.CallOption) (*CompositeType, error)
}

type compositeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompositeServiceClient(cc grpc.ClientConnInterface) CompositeServiceClient {
	return &compositeServiceClient{cc}
}

func (c *compositeServiceClient) UnaryExample(ctx context.Context, in *CompositeType, opts ...grpc.CallOption) (*CompositeType, error) {
	out := new(CompositeType)
	err := c.cc.Invoke(ctx, CompositeService_UnaryExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompositeServiceServer is the server API for CompositeService service.
// All implementations must embed UnimplementedCompositeServiceServer
// for forward compatibility
type CompositeServiceServer interface {
	UnaryExample(context.Context, *CompositeType) (*CompositeType, error)
	mustEmbedUnimplementedCompositeServiceServer()
}

// UnimplementedCompositeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompositeServiceServer struct {
}

func (UnimplementedCompositeServiceServer) UnaryExample(context.Context, *CompositeType) (*CompositeType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryExample not implemented")
}
func (UnimplementedCompositeServiceServer) mustEmbedUnimplementedCompositeServiceServer() {}

// UnsafeCompositeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompositeServiceServer will
// result in compilation errors.
type UnsafeCompositeServiceServer interface {
	mustEmbedUnimplementedCompositeServiceServer()
}

func RegisterCompositeServiceServer(s grpc.ServiceRegistrar, srv CompositeServiceServer) {
	s.RegisterService(&CompositeService_ServiceDesc, srv)
}

func _CompositeService_UnaryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompositeType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompositeServiceServer).UnaryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompositeService_UnaryExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompositeServiceServer).UnaryExample(ctx, req.(*CompositeType))
	}
	return interceptor(ctx, in, info, handler)
}

// CompositeService_ServiceDesc is the grpc.ServiceDesc for CompositeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompositeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "composite.CompositeService",
	HandlerType: (*CompositeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryExample",
			Handler:    _CompositeService_UnaryExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "composite/example.proto",
}