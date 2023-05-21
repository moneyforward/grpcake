package grpcake

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

// NewGrpcClient ...
func NewGrpcClient(ctx context.Context, url string, descSource DescriptorSource) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	client := grpcdynamic.NewStub(conn)

	return &GrpcClient{descriptorSource: descSource, client: client}, nil
}

// NewGrpcClientFromReflection ...
func NewGrpcClientFromReflection(ctx context.Context, url string) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	refClient := grpcreflect.NewClientV1Alpha(ctx, reflectpb.NewServerReflectionClient(conn))
	reflSource := DescriptorSourceFromServer(ctx, refClient)
	descSource := reflSource

	client := grpcdynamic.NewStub(conn)

	return &GrpcClient{
		descriptorSource: descSource,
		client:           client,
	}, nil
}

// NewGrpcClientFromProtoFiles ...
func NewGrpcClientFromProtoFiles(ctx context.Context, url string, importPaths, fileNames []string) (*GrpcClient, error) {
	if len(fileNames) == 0 {
		return nil, fmt.Errorf("error received empty list of files")
	}

	fileSource, err := DescriptorSourceFromProtoFiles(ctx, importPaths, fileNames)
	if err != nil {
		return nil, fmt.Errorf("failed to process proto source files: %s", err)
	}

	return NewGrpcClient(ctx, url, fileSource)
}

// InvokeRpc ...
func (g *GrpcClient) InvokeRpc(ctx context.Context, serviceName, methodName, jsonBody string, rawHeaders map[string]string) (proto.Message, error) {
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

	headers := metadata.New(rawHeaders)
	outgoingCtx := metadata.NewOutgoingContext(ctx, headers)

	resMsg, err := g.client.InvokeRpc(outgoingCtx, methodDescriptor, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error sending grpc request: %v", err)
	}

	return resMsg, nil
}
