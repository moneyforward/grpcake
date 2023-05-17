package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/linker"
	"github.com/tidwall/sjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

const (
	DefaultTimeout    = 15 * time.Second
	FilePathSeparator = ","
)

func main() {
	var (
		url        = flag.String("url", "", "GRPC Server URL")
		grpcMethod = flag.String("grpc-method", "", "GRPC Method")
		proto      = flag.String("proto", "", "Proto files to import")
	)

	flag.Parse()

	// parse the request body from non-flag arguments
	jsonBody, err := parseJSONFieldArg(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing json field arguments: %s", err)
		os.Exit(1)
	}

	fmt.Printf("request json body:\n %s\n", jsonBody)

	// get service URL and gRPC method
	if *url == "" || *grpcMethod == "" {
		fmt.Fprint(os.Stderr, "error url or grpc method is not passed")
		fmt.Fprintln(os.Stderr)
		os.Exit(2)
	}
	parts := strings.SplitN(*grpcMethod, "/", 2)
	if len(parts) != 2 {
		fmt.Fprint(os.Stderr, "error invalid grpc method name")
	}
	serviceName := parts[0]
	methodName := parts[1]

	fmt.Printf("service: %s \nmethod: %s\n", serviceName, methodName)

	// make client
	if *proto == "" {
		fmt.Fprint(os.Stderr, "error proto file is not passed")
		fmt.Fprintln(os.Stderr)
		os.Exit(2)
	}

	// construct client
	ctx := context.Background()

	grpcClient, err := NewGrpcClientFromProtoFiles(ctx, *url, *proto)
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}

	// send request
	timeoutCtx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	resProtoMsg, err := grpcClient.Send(timeoutCtx, serviceName, methodName, jsonBody)
	if err != nil {
		log.Fatalf("error sending grpc request: %v", err)
	}

	// convert proto msg response to bytes of json
	resMsgJSON, err := protojson.Marshal(resProtoMsg)
	if err != nil {
		log.Fatalf("error printing response message as JSON: %v", err)
	}

	fmt.Printf("Response:\n%s", resMsgJSON)
}

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
	return nil
}

// invokeRPC calls unary RPC methods on the server.
func (g GrpcClient) invokeRPC(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	return nil, nil
}

// parseJSONFieldArg Parse JSON field arguments into a json string.
func parseJSONFieldArg(args []string) (jsonString string, err error) {
	jsonString = "{}"

	var parts []string
	for _, arg := range args {
		parts = strings.SplitN(arg, ":=", 2)
		if len(parts) == 2 {
			jsonString, err = sjson.SetRaw(jsonString, parts[0], parts[1])
			if err != nil {
				return "", fmt.Errorf("error setting raw key value for json (%v, %v): %v", parts[0], parts[1], err)
			}
			continue
		}

		parts = strings.SplitN(arg, "=", 2)
		if len(parts) < 2 {
			return "", fmt.Errorf("error invalid format for arg '%v'", arg)
		}

		jsonString, err = sjson.Set(jsonString, parts[0], parts[1])
		if err != nil {
			return "", fmt.Errorf("error setting key value for json (%v, %v): %v", parts[0], parts[1], err)
		}
	}

	return jsonString, nil
}
