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
)

// GrpcClient invokes grpc method on a remote server dynamically, without the need for
// protobuf code generation.
type GrpcClient struct {
	fileDescriptors linker.Files
	client          grpc.ClientConnInterface
}

func NewGrpcClientFromProtoFiles(ctx context.Context, url string, protoFilePath string) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	compiler := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{}),
	}

	files, err := compiler.Compile(ctx, protoFilePath)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{fileDescriptors: files, client: conn}, nil
}

func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	// get service descriptor
	serviceDescriptor := getServiceDescriptorByFqnName(g.fileDescriptors, protoreflect.FullName(serviceName))
	if serviceDescriptor == nil {
		return nil, fmt.Errorf("error finding service with name %s", serviceName)
	}

	// get method descriptor
	methodDescriptor := serviceDescriptor.Methods().ByName(protoreflect.Name(methodName))
	if methodDescriptor == nil {
		return nil, fmt.Errorf("error finding method with name %s", methodName)
	}

	// create the proto msg
	reqMsg := dynamicpb.NewMessage(methodDescriptor.Input())
	err := protojson.Unmarshal([]byte(jsonBody), reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json body to protobuf message: %v", err)
	}

	// invoke the RPC method with the created msg
	resMsg, err := g.invokeRPC(ctx, methodDescriptor, reqMsg)
	if err != nil {
		return nil, fmt.Errorf("error sending grpc request: %v", err)
	}

	return resMsg, nil
}

// getServiceDescriptorByFqnName finds a service descriptor given a set of file descriptors.
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

// invokeRPC calls unary RPC methods on the server.
func (g GrpcClient) invokeRPC(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	if method.IsStreamingClient() || method.IsStreamingServer() {
		return nil, fmt.Errorf("InvokeRpc is for unary methods; %q is %s", method.FullName(), methodType(method))
	}

	// check msg type to make sure it matches what the method expects
	if err := checkMessageType(method.Input(), request); err != nil {
		return nil, fmt.Errorf("error checking message type: %v", err)
	}

	// make the gRPC call
	resp := dynamicpb.NewMessage(method.Output())
	if err := g.client.Invoke(ctx, requestMethod(method), request, resp, opts...); err != nil {
		return nil, fmt.Errorf("error invoking rpc method: %v", err)
	}

	return resp, nil
}

// checkMessageType checks if a given proto message fit with the given protoreflect.MessageDescriptor.
func checkMessageType(md protoreflect.MessageDescriptor, msg proto.Message) error {
	expectedMessageDescriptorFullName := md.FullName()
	givenMessageDescriptorFullName := msg.ProtoReflect().Descriptor().FullName()
	if expectedMessageDescriptorFullName != givenMessageDescriptorFullName {
		return fmt.Errorf(
			"error wrong message type: expecting %s, got %s",
			expectedMessageDescriptorFullName,
			givenMessageDescriptorFullName,
		)
	}

	return nil
}

// requestMethod generate method name string for invoking rpc methods.
func requestMethod(md protoreflect.MethodDescriptor) string {
	return fmt.Sprintf("/%s/%s", md.Parent().FullName(), md.Name())
}

// methodType returns a string to specify whether a method
// is unary, client streaming, server streaming or bidirectional streaming.
func methodType(md protoreflect.MethodDescriptor) string {
	if md.IsStreamingClient() && md.IsStreamingServer() {
		return "bidi-streaming"
	} else if md.IsStreamingClient() {
		return "client-streaming"
	} else if md.IsStreamingServer() {
		return "server-streaming"
	} else {
		return "unary"
	}
}
