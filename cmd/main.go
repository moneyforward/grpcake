package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"google.golang.org/grpc"
)

func main() {
	grpcClient, err := NewGrpcClientFromProtoFile("localhost:8080", "cmd/adder.proto")
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}

	resMsg, err := grpcClient.Send(context.Background(), "api.Adder", "Add", `{"x": 1, "y": 2}`)
	if err != nil {
		log.Fatalf("error sending grpc request: %v", err)
	}

	marshaler := jsonpb.Marshaler{}
	log.Println(marshaler.MarshalToString(resMsg))
}

type GrpcClient struct {
	fileDesc *desc.FileDescriptor
	conn     *grpc.ClientConn
	client   grpcdynamic.Stub
}

func NewGrpcClientFromProtoFile(url string, fileName string) (*GrpcClient, error) {
	// TODO: allow more options
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	client := grpcdynamic.NewStub(conn)

	parser := protoparse.Parser{}
	fileDescriptors, err := parser.ParseFiles(fileName)
	if err != nil {
		return nil, fmt.Errorf("error parsing proto file: %v", err)
	}

	if len(fileDescriptors) == 0 {
		return nil, fmt.Errorf("error no files found: %v", err)
	}

	return &GrpcClient{
		fileDesc: fileDescriptors[0],
		conn:     conn,
		client:   client,
	}, nil
}

func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	serviceDesc := g.fileDesc.FindService(serviceName)
	if serviceDesc == nil {
		return nil, fmt.Errorf("error service '%v' not found", serviceName)
	}

	methodDesc := serviceDesc.FindMethodByName(methodName)
	if methodDesc == nil {
		return nil, fmt.Errorf("error method '%v' not found in service '%v'", methodName, serviceName)
	}

	// create request protobuf message
	reqMsgDesc := methodDesc.GetInputType()
	if reqMsgDesc == nil {
		return nil, errors.New("todo")
	}
	reqMsg := dynamic.NewMessage(reqMsgDesc)

	// send grpc request
	err := jsonpb.Unmarshal(strings.NewReader(jsonBody), reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json body to protobuf message: %v", err)
	}

	resMsg, err := g.client.InvokeRpc(context.Background(), methodDesc, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error sending grpc request: %v", err)
	}

	return resMsg, nil
}
