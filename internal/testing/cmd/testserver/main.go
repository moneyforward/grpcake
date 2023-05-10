package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	grpcreflect "google.golang.org/grpc/reflection"

	testingpb "github.com/moneyforward/grpcake/internal/testing"
)

func main() {
	reflection := flag.Bool("use-reflection", false, "Run server with reflection support")
	port := 6069
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	testingpb.RegisterExampleServiceServer(grpcServer, newExampleServer())

	flag.Parse()

	if *reflection {
		// Register reflection service on gRPC server.
		grpcreflect.Register(grpcServer)
	}

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
