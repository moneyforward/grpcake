package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tidwall/sjson"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moneyforward/grpcake"
)

const (
	DefaultTimeout    = 15 * time.Second
	FilePathSeparator = ","
)

func main() {
	var (
		url            = flag.String("url", "", "GRPC Server URL")
		grpcMethod     = flag.String("grpc-method", "", "GRPC Method")
		rawProtoFiles  = flag.String("proto", "", "Proto files to import delimited by comma")
		rawImportPaths = flag.String("import-paths", "", "List of import paths delimited by comma")
	)

	flag.Parse()

	importPaths := make([]string, 0)
	if *rawImportPaths != "" {
		importPaths = strings.Split(*rawImportPaths, FilePathSeparator)
	}

	protoFiles := make([]string, 0)
	if *rawProtoFiles != "" {
		protoFiles = strings.Split(*rawProtoFiles, FilePathSeparator)
	}

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

	// construct client
	ctx := context.Background()

	var grpcClient *grpcake.GrpcClient
	if len(protoFiles) > 0 {
		// if proto files are given use them
		grpcClient, err = grpcake.NewGrpcClientFromProtoFiles(ctx, *url, protoFiles, importPaths)
	} else {
		// otherwise, use reflection
		grpcClient, err = grpcake.NewGrpcClientFromReflection(ctx, *url)
	}
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

	prettyResponse, err := JSONPrettify(resMsgJSON)
	if err != nil {
		log.Fatalf("error prettify-ing response json: %v", err)
	}
	fmt.Printf("Response:\n%s\n", prettyResponse)
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

// JSONPrettify adds indent to raw json string.
func JSONPrettify(jsonBytes []byte) (string, error) {
	var prettifiedJSON bytes.Buffer
	err := json.Indent(&prettifiedJSON, jsonBytes, "", "\t")
	if err != nil {
		return "", fmt.Errorf("error prettify-ing jsong string: %v", err)
	}

	return prettifiedJSON.String(), nil
}
