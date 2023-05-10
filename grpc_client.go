package grpcake

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/moneyforward/grpcake/internal/grpcdynamic"
)

// GrpcClient invokes grpc method on a remote server dynamically, without the need for
// protobuf code generation.
type GrpcClient struct {
	descriptorSource DescriptorSource
	client           grpcdynamic.Stub
}

// NewGrpcClientFromProtoFiles ...
func NewGrpcClientFromProtoFiles(url string, fileNames []string) (*GrpcClient, error) {

	ctx := context.Background()

	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	client := grpcdynamic.NewStub(conn)

	fileSource, err := DescriptorSourceFromProtoFiles(ctx, fileNames...)
	if err != nil {
		return nil, fmt.Errorf("failed to process proto source files: %s", err)
	}

	return &GrpcClient{descriptorSource: fileSource, client: client}, nil
}

func NewGrpcClientFromReflectingServer(url string) (*GrpcClient, error) {
	ctx := context.Background()

	// TODO: remove code duplication
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}
	client := grpcdynamic.NewStub(cc)

	refClient := grpcreflect.NewClientV1Alpha(ctx, reflectpb.NewServerReflectionClient(cc))
	reflSource := DescriptorSourceFromServer(ctx, refClient)
	return &GrpcClient{descriptorSource: reflSource, client: client}, nil
}

// Send ...
func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	serviceDescriptor, err := g.descriptorSource.FindServiceDescriptor(serviceName)
	if serviceDescriptor == nil || err != nil {
		return nil, fmt.Errorf("error finding service with name %s: %s", serviceName, err)
	}

	methodDescriptor := serviceDescriptor.Methods().ByName(protoreflect.Name(methodName))
	if methodDescriptor == nil {
		return nil, fmt.Errorf("error finding method with name %s", methodName)
	}

	reqMsg := dynamicpb.NewMessage(methodDescriptor.Input())

	err = protojson.Unmarshal([]byte(jsonBody), reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json body to protobuf message: %v", err)
	}

	// TODO:
	resMsg, err := g.client.InvokeRpc(context.Background(), methodDescriptor, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error sending grpc request: %v", err)
	}

	return resMsg, nil
}
