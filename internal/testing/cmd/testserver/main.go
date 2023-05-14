package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	grpcreflect "google.golang.org/grpc/reflection"

	barpb "github.com/moneyforward/grpcake/internal/testing/pb/bar"
	bazpb "github.com/moneyforward/grpcake/internal/testing/pb/baz"
	foopb "github.com/moneyforward/grpcake/internal/testing/pb/foo"
)

func main() {
	reflection := flag.Bool("use-reflection", false, "Run server with reflection support")

	flag.Parse()

	port := 6069
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	foopb.RegisterExampleServiceServer(grpcServer, newFooServer())
	barpb.RegisterTestServiceServer(grpcServer, newBarServer())
	bazpb.RegisterTestServiceServer(grpcServer, newBazServer())

	if *reflection {
		// Register reflection service on gRPC server.
		grpcreflect.Register(grpcServer)
	}

	log.Printf("Service is running on localhost:%d", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}

type fooServer struct {
	*foopb.UnimplementedExampleServiceServer
}

func newFooServer() *fooServer {
	return &fooServer{}
}

func (e *fooServer) UnaryExample(ctx context.Context, pb *foopb.BasicTypes) (*foopb.BasicTypes, error) {
	return pb, nil
}

type barServer struct {
	*barpb.UnimplementedTestServiceServer
}

func newBarServer() *barServer {
	return &barServer{}
}

func (b *barServer) UnaryExample(ctx context.Context, pb *barpb.ExampleMessage) (*barpb.ExampleMessage, error) {
	return pb, nil
}

type bazServer struct {
	*bazpb.UnimplementedTestServiceServer
}

func newBazServer() *bazServer {
	return &bazServer{}
}

func (b *bazServer) Greeting(ctx context.Context, pb *bazpb.GreetingMessage) (*bazpb.GreetingMessage, error) {
	return pb, nil
}
