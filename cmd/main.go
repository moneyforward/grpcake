package main

import (
	"context"
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

// TODO: write tests
func main() {
	var (
		url            = flag.String("url", "", "GRPC Server URL")
		reflection     = flag.Bool("use-reflection", true, "Try to use server reflection")
		grpcMethod     = flag.String("grpc-method", "", "GRPC Method")
		rawImportFiles = flag.String("proto", "", "Proto files to import")
		jsonBody       = flag.String("body", "", "JSON body")
		rawImportPaths = flag.String("import-path", "", "List of import paths")
	)

	flag.Parse()

	if *jsonBody == "" {
		jsonString, err := parseJSONFieldArg(flag.Args())
		if err != nil {
			fail(err, "error parsing json field arguments")
		}

		jsonBody = &jsonString
	}

	var importPaths []string
	if *rawImportPaths != "" {
		importPaths = strings.Split(*rawImportPaths, FilePathSeparator)
	}

	log.Printf("request json body: %s", *jsonBody)

	if *url == "" || *grpcMethod == "" {
		fail(nil, "error url or grpc method is not passed")
	}

	if !*reflection && *rawImportFiles == "" {
		fail(nil, "No protoset files or proto files specified and -use-reflection set to false.")
	}

	globalCtx := context.Background()

	importFiles := make([]string, 0)
	if *rawImportFiles != "" {
		importFiles = strings.Split(*rawImportFiles, FilePathSeparator)
	}

	var grpcClient *grpcake.GrpcClient
	var err error
	if len(importFiles) > 0 {
		grpcClient, err = grpcake.NewGrpcClientFromProtoFiles(globalCtx, *url, importPaths, importFiles)
		if err != nil {
			log.Fatalf("error creating grpc client: %v", err)
		}
	} else {
		// Fallback to server reflections for proto descriptor discovery
		grpcClient, err = grpcake.NewGrpcClientFromReflection(globalCtx, *url)
		if err != nil {
			log.Fatalf("error creating grpc client: %v", err)
		}
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

	resMsgJSON, err := protojson.Marshal(resMsg)
	if err != nil {
		log.Fatalf("error printing response message as JSON: %v", err)
	}

	prettiedJSON, err := grpcake.JSONPrettify(resMsgJSON)
	if err != nil {
		log.Fatalf("error prettify-ing response json: %v", err)
	}
	log.Println("Response:", prettiedJSON)
}

// parseJSONFieldArg Parse JSON field arguments into a json string.
// take a look at this case "e f:"=a will be parsed as e f:=a
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

func fail(err error, msg string, args ...interface{}) {
	if err != nil {
		msg += ": %v"
		args = append(args, err)
	}
	fmt.Fprintf(os.Stderr, msg, args...)
	fmt.Fprintln(os.Stderr)
	if err != nil {
		os.Exit(1)
	} else {
		// nil error means it was CLI usage issue
		fmt.Fprintf(os.Stderr, "Try '%s -help' for more details.\n", os.Args[0])
		os.Exit(2)
	}
}
