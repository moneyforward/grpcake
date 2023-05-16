package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/sjson"
)

func main() {
	var (
		grpcMethod = flag.String("grpc-method", "", "GRPC Method")
	)

	flag.Parse()

	// parse the request body from non-flag arguments
	jsonBody, err := parseJSONFieldArg(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing json field arguments: %s", err)
		os.Exit(1)
	}

	fmt.Printf("request json body:\n %s\n", jsonBody)

	parts := strings.SplitN(*grpcMethod, "/", 2)
	if len(parts) != 2 {
		fmt.Fprint(os.Stderr, "error invalid grpc method name")
	}

	serviceName := parts[0]
	methodName := parts[1]

	fmt.Printf("service: %s \nmethod: %s\n", serviceName, methodName)
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
