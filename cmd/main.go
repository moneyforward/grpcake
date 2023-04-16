package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/tidwall/sjson"
	"log"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"google.golang.org/grpc"
)

const (
	DefaultTimeout = 15 * time.Second
)

// TODO: write tests
func main() {
	url := flag.String("url", "", "GRPC Server URL")
	grpcMethod := flag.String("grpc-method", "", "GRPC Method")
	importFileName := flag.String("import", "", "Proto files to import")
	jsonBody := flag.String("body", "", "JSON body")

	flag.Parse()

	if *jsonBody == "" {
		jsonString, err := parseJsonFieldArg(flag.Args())
		if err != nil {
			log.Fatalf("error parsing json field arguments: %v", err)
		}

		jsonBody = &jsonString
	}

	log.Println("request json body: %s", *jsonBody)

	if *url == "" || *grpcMethod == "" {
		log.Fatalf("error url or grpc method is not passed")
	}

	globalCtx := context.Background()
	grpcClient, err := NewGrpcClientFromProtoFile(*url, *importFileName)
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}

	timeoutCtx, cancel := context.WithTimeout(globalCtx, DefaultTimeout)
	defer cancel()

	parts := strings.SplitN(*grpcMethod, "/", 2)
	if len(parts) != 2 {
		log.Fatalf("error invalid grpc method name")
	}

	serviceName := parts[0]
	methodName := parts[1]

	resMsg, err := grpcClient.Send(timeoutCtx, serviceName, methodName, *jsonBody)
	if err != nil {
		log.Fatalf("error sending grpc request: %v", err)
	}

	marshaler := jsonpb.Marshaler{}
	resMsgJson, err := marshaler.MarshalToString(resMsg)
	if err != nil {
		log.Fatalf("error printing response message as JSON: %v", err)
	}

	log.Println(resMsgJson)
}

// parseJsonFieldArg Parse JSON field arguments into a json string.
// take a look at this case "e f:"=a will be parsed as e f:=a
func parseJsonFieldArg(args []string) (jsonString string, err error) {
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

// GrpcClient invokes grpc method on a remote server dynamically, without the need for
// protobuf code generation.
type GrpcClient struct {
	fileDesc *desc.FileDescriptor
	conn     *grpc.ClientConn
	client   grpcdynamic.Stub
}

// TODO: support importing multiple files
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

// Send ...
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
