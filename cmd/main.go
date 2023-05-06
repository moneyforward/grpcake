package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tidwall/sjson"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/moneyforward/grpcake"
)

const (
	DefaultTimeout = 15 * time.Second
)

// TODO: write tests
func main() {
	url := flag.String("url", "", "GRPC Server URL")
	grpcMethod := flag.String("grpc-method", "", "GRPC Method")
	importFileNames := flag.String("import", "", "Proto files to import")
	jsonBody := flag.String("body", "", "JSON body")

	flag.Parse()

	if *jsonBody == "" {
		jsonString, err := parseJsonFieldArg(flag.Args())
		if err != nil {
			log.Fatalf("error parsing json field arguments: %v", err)
		}

		jsonBody = &jsonString
	}

	log.Printf("request json body: %s", *jsonBody)

	if *url == "" || *grpcMethod == "" {
		log.Fatalf("error url or grpc method is not passed")
	}

	globalCtx := context.Background()
	grpcClient, err := grpcake.NewGrpcClientFromProtoFiles(*url, strings.Split(*importFileNames, " "))
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

	resMsgJson, err := protojson.Marshal(resMsg)
	if err != nil {
		log.Fatalf("error printing response message as JSON: %v", err)
	}

	prettifiedJson, err := grpcake.JsonPrettify(resMsgJson)
	if err != nil {
		log.Fatalf("error prettify-ing response json: %v", err)
	}
	log.Println("Response:", prettifiedJson)
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
