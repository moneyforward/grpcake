package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	advancedpb "github.com/moneyforward/grpcake/internal/testing/pb/advanced"
	basicpb "github.com/moneyforward/grpcake/internal/testing/pb/basic"
	compositepb "github.com/moneyforward/grpcake/internal/testing/pb/composite"
)

func main() {
	port := 6069
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor((&loggingInterceptor{}).UnaryServerInterceptor),
	)
	basicpb.RegisterBasicServiceServer(grpcServer, newBasicServer())
	advancedpb.RegisterAdvancedServiceServer(grpcServer, newAdvancedServer())
	compositepb.RegisterCompositeServiceServer(grpcServer, newCompositeServer())

	log.Printf("Service is running on localhost:%d", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}

type basicServer struct {
	*basicpb.UnimplementedBasicServiceServer
}

func newBasicServer() *basicServer {
	return &basicServer{}
}

func (e *basicServer) UnaryExample(ctx context.Context, pb *basicpb.BasicType) (*basicpb.BasicType, error) {
	return pb, nil
}

type advancedServer struct {
	*advancedpb.UnimplementedAdvancedServiceServer
}

func newAdvancedServer() *advancedServer {
	return &advancedServer{}
}

func (b *advancedServer) UnaryExample(ctx context.Context, pb *advancedpb.AdvancedType) (*advancedpb.AdvancedType, error) {
	return pb, nil
}

type compositeServer struct {
	*compositepb.UnimplementedCompositeServiceServer
}

func newCompositeServer() *compositeServer {
	return &compositeServer{}
}

func (b *compositeServer) UnaryExample(ctx context.Context, pb *compositepb.CompositeType) (*compositepb.CompositeType, error) {
	return pb, nil
}

type loggingInterceptor struct{}

func (i *loggingInterceptor) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Log the method name
	log.Printf("gRPC method called: %s", info.FullMethod)

	// Log the request headers
	// md, _ := metadata.FromIncomingContext(ctx)
	// log.Printf("Request headers: %v", md)

	// Call the next handler in the chain
	return handler(ctx, req)
}
