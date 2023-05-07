package grpcake

import (
	"context"
	"fmt"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/linker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/moneyforward/grpcake/internal/grpcdynamic"
)

// GrpcClient invokes grpc method on a remote server dynamically, without the need for
// protobuf code generation.
type GrpcClient struct {
	fileDescriptors linker.Files
	client          grpcdynamic.Stub
}

// NewGrpcClientFromProtoFiles ...
func NewGrpcClientFromProtoFiles(url string, fileNames []string) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	client := grpcdynamic.NewStub(conn)

	compiler := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{}),
	}

	files, err := compiler.Compile(context.Background(), fileNames...)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{fileDescriptors: files, client: client}, nil
}

// Send ...
func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	serviceDescriptor := getServiceDescriptorByFqnName(g.fileDescriptors, protoreflect.FullName(serviceName))
	if serviceDescriptor == nil {
		return nil, fmt.Errorf("error finding service with name %s", serviceName)
	}

	methodDescriptor := serviceDescriptor.Methods().ByName(protoreflect.Name(methodName))
	if methodDescriptor == nil {
		return nil, fmt.Errorf("error finding method with name %s", methodName)
	}

	reqMsg := dynamicpb.NewMessage(methodDescriptor.Input())

	err := protojson.Unmarshal([]byte(jsonBody), reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json body to protobuf message: %v", err)
	}

	resMsg, err := g.client.InvokeRpc(context.Background(), methodDescriptor, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error sending grpc request: %v", err)
	}

	return resMsg, nil
}

// getServiceDescriptorByFqnName ...
func getServiceDescriptorByFqnName(fileDescriptors linker.Files, serviceName protoreflect.FullName) protoreflect.ServiceDescriptor {
	for _, descriptor := range fileDescriptors {
		svcDescriptors := descriptor.Services()
		for i := 0; i < svcDescriptors.Len(); i++ {
			serviceDescriptor := svcDescriptors.Get(i)
			if serviceDescriptor.FullName() == serviceName {
				return serviceDescriptor
			}
		}
	}

	return nil
}
