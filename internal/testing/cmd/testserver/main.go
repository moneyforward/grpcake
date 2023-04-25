package main

import (
	"context"
	"fmt"
	"log"
	"net"

	testingpb "github.com/moneyforward/grpcake/internal/testing"
	"google.golang.org/grpc"
)

func main() {
	port := 6069
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	testingpb.RegisterExampleServiceServer(grpcServer, newExampleServer())

	log.Printf("Service is running on localhost:%d", port)
	grpcServer.Serve(lis)
}

// TODO: add the rest of the methods
type exampleServer struct {
	*testingpb.UnimplementedExampleServiceServer
}

func newExampleServer() *exampleServer {
	return &exampleServer{}
}

func (e *exampleServer) UnaryExample(ctx context.Context, pb *testingpb.BasicTypes) (*testingpb.BasicTypes, error) {
	return pb, nil
}
