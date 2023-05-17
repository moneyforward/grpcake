# Hands on proto reflect v2

*TODO: find explaination for all the "explain ..." parts*

## Intro

### What is gRPC

### How a gRPC client works

### The problem we are trying to solve

### What will learn doing that

## PART I: Import single proto file

### Init

We will start from commit [7175b14](https://github.com/moneyforward/grpcake/tree/7175b14fa588d7b80e8c4662068b1e738b7189c2).

It only has the gRPC service for testing.

### Read gRPC service and method and parse arguments as JSON

[d006516](https://github.com/moneyforward/grpcake/commit/d0065165e6fcb6c59f964b34968497a27de45d00)

- create a `cmd/main.go` as follow
- explain that with `flag.Args()` we get all the non-flag arguments which we will use to construct the
  parameters to the gRPC request.

```go
// cmd/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		grpcMethod = flag.String("grpc-method", "", "GRPC Method")
	)

	flag.Parse()

	// parse the request body from non-flag arguments
	jsonBody := parseJSONFieldArg(flag.Args())

	fmt.Printf("request json body:\n %s\n", jsonBody)

	parts := strings.SplitN(*grpcMethod, "/", 2)
	if len(parts) != 2 {
		fmt.Fprint(os.Stderr, "error invalid grpc method name")
	}

	serviceName := parts[0]
	methodName := parts[1]

	fmt.Printf("service: %s \nmethod: %s\n", serviceName, methodName)

	// construct client
	// make request
	// print response
}

func parseJSONFieldArg(args []string) string {
	return ""
}
```

#### Add body of `parseJSONFieldArg`

- Explain the idea behind `=` and `:=`.
- Explain the differences between functions is `tidwall/sjson` we will use
  - `SetRaw` 
  - `Set`

```go
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
```

Update imports with (the editor should do it on save)

```go
import "github.com/tidwall/sjson"
```


### Get service URL

[6caba19](https://github.com/moneyforward/grpcake/commit/6caba19bb8ba140a913fd80dc95ed9edace7a59c)

- add the flag for reading the url
- check that url and gRPC method are set
- if not exit

### Read path to proto file

[d8829f4](https://github.com/moneyforward/grpcake/commit/d8829f4fda9f75856d3ca1314f8567da9626cfd6)

- add flag for reading one proto file
- if not set exit

### 1b7330b Construct client

- update the part under the comment that says construct client
- add the `GrpcClient` struct
  - introduce the `protocompile` package
  - explain what is `linker.Files` 
  - explain what is `grpc.ClientConnInterface`
- add the definition of `NewGrpcClientFromProtoFiles`


```go
func main() {
   . . .
   // construct client
	grpcClient, err := NewGrpcClientFromProtoFiles(ctx, *url, *proto)
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}
}

// GrpcClient invokes grpc method on a remote server dynamically, without the need for
// protobuf code generation.
type GrpcClient struct {
	fileDescriptors linker.Files
	client          grpc.ClientConnInterface
}

func NewGrpcClientFromProtoFiles(ctx context.Context, url string, protoFilePath string) (*GrpcClient, error) {
	return &GrpcClient{nil, nil}, nil
}
```



### Implement helper for constructing client

[183d376](https://github.com/moneyforward/grpcake/commit/183d376bc6410794ee2b2320810672bc164b0798)

- explain what is `grpc.Dial`
- explain what is `protocompile.Compiler`

```go
func NewGrpcClientFromProtoFiles(ctx context.Context, url string, protoFilePath string) (*GrpcClient, error) {
	return &GrpcClient{nil, nil}, nil
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
```


### Send request and print converted response

[48b057a](https://github.com/moneyforward/grpcake/commit/48b057a7a5738afa20c0abeec74201761889b880)

- update the part under the comment that says "send request" in cmd/main.go

```go
func main() {
	...
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
```

- add definition `GrpcClient.Send`

```go
func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	return nil, nil
}
```

### Add steps for send method

[17315de](https://github.com/moneyforward/grpcake/commit/17315de5e8ec17c10277fff03a04758360eaa3b2)

- explain the steps for implementing `GrpcClient.Send`
- show the documentation of descriptors
- explain the concept of descriptor names and full names

```go
func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	return nil, nil
	// get service descriptor from file descriptors
	// get method descriptor from service descriptor
	// create the proto msg
	// invoke the RPC method with the created msg
}
```

### Implement body of Send and add helper declaration of helper methods

[4f11a31](https://github.com/moneyforward/grpcake/commit/4f11a3151d62193ad60b0c28d075d60c2763715a)

- implement body of send
- explain what `dyanamicpb.NewMessage` does

```go
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
```

- add declaration of helper functions and methods

```go
// getServiceDescriptorByFqnName finds a service descriptor given a set of file descriptors.
func getServiceDescriptorByFqnName(fileDescriptors linker.Files, serviceName protoreflect.FullName) protoreflect.ServiceDescriptor {
	return nil
}

// invokeRPC calls unary RPC methods on the server.
func (g GrpcClient) invokeRPC(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	return nil, nil
}
```

### Implement helper for finding service descriptor

[28445c4](https://github.com/moneyforward/grpcake/commit/28445c43884b252cd608837a99060f1a88aba57f)

```go
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
```

### Add steps for body of invokeRPC

[86a14b4](https://github.com/moneyforward/grpcake/commit/86a14b442d0c5420a5cb6d6af1c091f4bec545da)

```go
// invokeRPC calls unary RPC methods on the server.
func (g GrpcClient) invokeRPC(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	return nil, nil
	// check msg type to make sure it matches what the method expects
	// make the gRPC call
}
```

### Implement body of invokeRPC and add declaration of helpers

[351e84d](https://github.com/moneyforward/grpcake/commit/351e84d8a1e77c3a85ec8b676bf9503897b69699)
 
- add body of `invokeRPC`
- stress the fact that `dynamicpb.NewMessage` is used for receiving the response of the gRPC call

```go
// invokeRPC calls unary RPC methods on the server.
func (g GrpcClient) invokeRPC(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	return nil, nil
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
```

- add declaration of helpers

```go
// checkMessageType checks if a given proto message fit with the given protoreflect.MessageDescriptor.
func checkMessageType(md protoreflect.MessageDescriptor, msg proto.Message) error {
	return nil
}

// requestMethod generate method name string for invoking rpc methods.
func requestMethod(md protoreflect.MethodDescriptor) string {
	return ""
}

// methodType returns a string to specify whether a method
// is unary, client streaming, server streaming or bidirectional streaming.
func methodType(md protoreflect.MethodDescriptor) string {
	return ""
}
```

### Implement body of checkMessageType

[ab3822c](https://github.com/moneyforward/grpcake/commit/ab3822c27368c1f78f9e668336bc2faec6af1171)

- explain why this step is necessary

```go
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
```

### Implement body of requestMethod

[1b7c079](https://github.com/moneyforward/grpcake/commit/1b7c0790e7ceff73bc2ed25c7ff5293e5a00cf0a)

```go
// requestMethod generate method name string for invoking rpc methods.
func requestMethod(md protoreflect.MethodDescriptor) string {
	return ""
	return fmt.Sprintf("/%s/%s", md.Parent().FullName(), md.Name())
}
```

### Complete send by implementing body of methodType

[fabcafd](https://github.com/moneyforward/grpcake/commit/fabcafdf231a0378ffb309a682357dde5b8e1459)

```go
func methodType(md protoreflect.MethodDescriptor) string {
	return ""
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
```

### Time to Try

```
 $  go run cmd/main.go -url localhost:6069 \
    -grpc-method foo.ExampleService/UnaryExample \
    -proto internal/testing/proto/foo/example.proto \
    long_field:=10 string_field="hello world"

 request json body:
  {"long_field":10,"string_field":"hello world"}
  service: foo.ExampleService
  method: UnaryExample
  Response:
  {"longField":"10","stringField":"hello world"}
```

### Put GrpcClient and helpers in separate file

[b93c0c5](https://github.com/moneyforward/grpcake/commit/b93c0c59008e04548b1d132aa784c7ea0d5b1645)

- this is a refactoring step
- move everything related to `GrpcClient` to `./grpc_client.go`
- update reference in `cmd/main.go` as needed
- add `import github.com/moneyforward/grpcake` in `cmd/main.go`

## PART 2: Implement support for reflection

- explain that we only need to query the server for the service descriptor

### Fall back to reflection when proto file is not passed

[d9c780e](https://github.com/moneyforward/grpcake/commit/d9c780e16826dfe720526ef12a7d8b6622c1525a) 

- In the function `main` in `cmd/main.go` remove this `if` block

  ```
  func main() {
  	...
  	if *proto == "" {
  		fmt.Fprint(os.Stderr, "error proto file is not passed")
  		fmt.Fprintln(os.Stderr)
  		os.Exit(2)
  	}
  	...
  }
  ```

- then replace the line 

  ```go
  	grpcClient, err := grpcake.NewGrpcClientFromProtoFiles(ctx, *url, *proto)
  ```

  with 
  ```go
	grpcClient, err := grpcake.NewGrpcClientFromProtoFiles(ctx, *url, *proto)
	var grpcClient *grpcake.GrpcClient
	if *proto != "" {
		// if proto file is given use it
		grpcClient, err = grpcake.NewGrpcClientFromProtoFiles(ctx, *url, *proto)
	} else {
		// otherwise, use reflection
		grpcClient, err = grpcake.NewGrpcClientFromReflection(ctx, *url)
	}
	if err != nil {
		log.Fatalf("error creating grpc client: %v", err)
	}
  ```
  
### Refactor by introducing DescriptorSource

[a292f6c](https://github.com/moneyforward/grpcake/commit/a292f6cdac8cf8ad041be665986669c7e543548d)

- this is a refactoring step in `grpc_client.go`

- change the definition of `GrpcClient` to

```go
type GrpcClient struct {
	descriptorSource DescriptorSource
	client           grpc.ClientConnInterface
}
```

- define the interface `DescriptorSource`
  - explain that because all we need is the service descriptor. So it makes sense to abstract that away.

```go
// DescriptorSource is a source of protobuf descriptor information. It can be backed by a FileDescriptorSet
// proto (like a file generated by protoc) or a remote server that supports the reflection API.
type DescriptorSource interface {
	// FindSymbol returns a descriptor for the given fully-qualified symbol name.
	FindServiceDescriptor(fullyQualifiedName string) (protoreflect.ServiceDescriptor, error)
}
```

- update `GrpcClient.Send` so that it calls `GrpcClient.descriptorSource.FindServiceDescriptor`

```go
func (g *GrpcClient) Send(ctx context.Context, serviceName, methodName, jsonBody string) (proto.Message, error) {
	// get service descriptor
	serviceDescriptor, err := g.descriptorSource.FindServiceDescriptor(serviceName)
	if err != nil {
		return nil, fmt.Errorf("error finding service with name %s: %v", serviceName, err)
	}
	...
}
```

- declare the `fileSource` struct

```go
type fileSource struct {
	files linker.Files
}
```

- make `fileSource` implement `DescriptorSource` by changing `getServiceDescriptorByFqnName` so that it becomes
  the `FindServiceDescriptor` implementation of `fileSource`

```go
// FindServiceDescriptor implements DescriptorSource
func (fs fileSource) FindServiceDescriptor(fullyQualifiedName string) (protoreflect.ServiceDescriptor, error) {
	for _, descriptor := range fs.files {
		svcDescriptors := descriptor.Services()
		for i := 0; i < svcDescriptors.Len(); i++ {
			serviceDescriptor := svcDescriptors.Get(i)
			if serviceDescriptor.FullName() == protoreflect.FullName(fullyQualifiedName) {
				return serviceDescriptor, nil
			}
		}
	}

	return nil, fmt.Errorf("error finding service with name %s", fullyQualifiedName)
}
```

- add a constructor for `fileSource`

```go
func DescriptorSourceFromProtoFiles(ctx context.Context, fileNames ...string) (DescriptorSource, error) {
	compiler := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{}),
	}
	files, err := compiler.Compile(ctx, fileNames...)
	if err != nil {
		return nil, err
	}
	return fileSource{files: files}, nil
}
```

- update the body of `NewGrpcClientFromProtoFiles` so that it calls `DescriptorSourceFromProtoFiles`

```go
func NewGrpcClientFromProtoFiles(ctx context.Context, url string, protoFilePath string) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	fileSource, err := DescriptorSourceFromProtoFiles(ctx, protoFilePath)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{descriptorSource: fileSource, client: conn}, nil
}
```

### Implement descriptor source from server reflection

[036daab](https://github.com/moneyforward/grpcake/commit/036daab49dcf7d1b49149e08a6708df1697348d4)

- add a `serverSource` struct
- explain that it will query the server for the service descriptor
- describe the `grpcreflect` package and explain what is `grpcreflect.Client`
  - explain that we can use `grpcreflect.Client.ResolveService` for requesting a service descriptor
    if the server supports reflection
  - explain that we can use `grpcreflect.Client.ResolveService` is a gRPC method that the server implements when it supports
    reflection

```go
type serverSource struct {
	client *grpcreflect.Client
}
```

- make `serverSource` implement `DescriptorSource`

```go
func (ss serverSource) FindServiceDescriptor(fullyQualifiedName string) (protoreflect.ServiceDescriptor, error) {
	sd, err := ss.client.ResolveService(fullyQualifiedName)
	if err != nil {
		if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
			return nil, errors.New("server does not support the reflection API")
		}
		return nil, err
	}
	return sd.UnwrapService(), nil
}
```

- explain that we can use `grpcreflect.Client.ResolveService` is a gRPC method that the server implements when it supports
  reflection

```go
// DescriptorSourceFromServer creates a DescriptorSource that uses the given gRPC reflection client
// to interrogate a server for descriptor information. If the server does not support the reflection
// API then the various DescriptorSource methods will return ErrReflectionNotSupported
func DescriptorSourceFromServer(_ context.Context, refClient *grpcreflect.Client) DescriptorSource {
	return serverSource{client: refClient}
}
```

### Implement body of `NewGrpcClientFromReflection`

[0d8a704](https://github.com/moneyforward/grpcake/commit/0d8a70443942577206cec3b8bd8d35c700df5515)

It was used in `cmd/main.go` to fall back on reflection 


- add this in `grpc_client.go`

```go
// NewGrpcClientFromReflection returns a GrpcClient that queries the server for service descriptors.
func NewGrpcClientFromReflection(ctx context.Context, url string) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	refClient := grpcreflect.NewClientV1Alpha(ctx, reflectpb.NewServerReflectionClient(conn))
	reflSource := DescriptorSourceFromServer(ctx, refClient)
	descSource := reflSource

	return &GrpcClient{
		descriptorSource: descSource,
		client:           conn,
	}, nil
}
```

- explain what is `grpcreflect.NewClientV1Alpha`

### Time to Try

- run server without reflection and call without proto file

```
$ make run-test-server
$ go run cmd/main.go -url localhost:6069 \
    -grpc-method foo.ExampleService/UnaryExample \
    long_field:=10 string_field="hello world"
```
expected response
```
 request json body:
  {"long_field":10,"string_field":"hello world"}
  service: foo.ExampleService
  method: UnaryExample
  error sending grpc request: error finding service with name foo.ExampleService: server does not support the reflection API
exit status 1
```

- run server with reflection and call without proto file

```
$ make run-test-server REFLECT=true
$ go run cmd/main.go -url localhost:6069 \
    -grpc-method foo.ExampleService/UnaryExample \
    long_field:=10 string_field="hello world"
```
expected response
```
request json body:
  {"long_field":10,"string_field":"hello world"}
service: foo.ExampleService
method: UnaryExample
Response:
{"longField":"10","stringField":"hello world"}
```

### Put DescriptorSource in separate file

[5e1079a](https://github.com/moneyforward/grpcake/commit/5e1079a51195d51f0d31bd4fb146363ef80d2466)

- this is a refactoring step
- move everything related to `DescriptorSource` in a file named `desc_source.go`
- no need update reference because they are still in the same package


## PART 3: Implement support for multiple proto files

- explain that proto files will be read separated by comma `,` and import paths also.
- explain the role of import paths

### update main

[ba0719a] 

- rename `proto      = flag.String("proto", "", "Proto files to import delimited by comma")` to `rawProtoFiles`
- add flag `rawImportPaths      = flag.String("import-paths", "", "List of import paths delimited by comma")` to `rawProtoFiles`

The flag block in main should look like this

```go
func main() {
	var (
		url            = flag.String("url", "", "GRPC Server URL")
		grpcMethod     = flag.String("grpc-method", "", "GRPC Method")
		rawProtoFiles  = flag.String("proto", "", "Proto files to import delimited by comma")
		rawImportPaths = flag.String("import-paths", "", "List of import paths delimited by comma")
	)
	
	flag.Parse()
	...
}
```

- make lists of import paths and proto files
```go
func main() {
	...
	flag.Parse()

	importPaths := make([]string, 0)
	if *rawImportPaths != "" {
		importPaths = strings.Split(*rawImportPaths, FilePathSeparator)
	}

	protoFiles := make([]string, 0)
	if *rawProtoFiles != "" {
		protoFiles = strings.Split(*rawProtoFiles, FilePathSeparator)
	}
	...
}
```
- in the `if else` block where the client are generated, update the condition for generating from proto files
  and update the call to `grpcake.NewGrpcClientFromProtoFiles`

```go
if len(protoFiles) > 0 {
		// if proto files are given use them
		grpcClient, err = grpcake.NewGrpcClientFromProtoFiles(ctx, *url, protoFiles, importPaths)
	} else {
		// otherwise, use reflection
		...
	}
```

### update definition of NewGrpcClientFromProtoFiles

[38d758c] 

- in `grpc_client.go`, update the definition of `NewGrpcClientFromProtoFiles` 
- in  the body of `NewGrpcClientFromProtoFiles` update the call to `DescriptorSourceFromProtoFiles`

```go
func NewGrpcClientFromProtoFiles(ctx context.Context, url string, protoFiles, importPaths []string) (*GrpcClient, error) {
	conn, err := dial(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to grpc server: %v", err)
	}

	fileSource, err := DescriptorSourceFromProtoFiles(ctx, protoFiles, importPaths)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{descriptorSource: fileSource, client: conn}, nil
}
```

### update definition of DescriptorSourceFromProtoFiles

[7157aa0] 

- update the definition of `DescriptorSourceFromProtoFiles`

```go
func DescriptorSourceFromProtoFiles(ctx context.Context, protoFiles, importPaths []string) (DescriptorSource, error) {
	compiler := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{
			ImportPaths: importPaths,
		}),
	}
	files, err := compiler.Compile(ctx, protoFiles...)
	if err != nil {
		return nil, err
	}
```

### Prettify JSON output (Optional)

[db84cca] 

- in `cmd/main.go` add

```go
// JSONPrettify adds indent to raw json string.
func JSONPrettify(jsonBytes []byte) (string, error) {
	var prettifiedJSON bytes.Buffer
	err := json.Indent(&prettifiedJSON, jsonBytes, "", "\t")
	if err != nil {
		return "", fmt.Errorf("error prettify-ing jsong string: %v", err)
	}

	return prettifiedJSON.String(), nil
}
```

- replace the last line of the `main` function   
  ```
	fmt.Printf("Response:\n%s", resMsgJSON)
  ```
  with
  ```go
  	prettyResponse, err := JSONPrettify(resMsgJSON)
  	if err != nil {
  		log.Fatalf("error prettify-ing response json: %v", err)
  	}
  	fmt.Printf("Response:\n%s\n", prettyResponse)
  ```

### Time To Try

```
$ make run-test-server
$ go run cmd/main.go --url localhost:6069 \
  --grpc-method bar.TestService/UnaryExample \
  --import-path internal/testing/proto \
  --proto bar/example.proto,foo/example.proto \
  int32Field:=1 basicTypes.intField:=2
```
expected response
```
request json body: 
 {"int32Field":1,"basicTypes":{"intField":2}}
service: bar.TestService
method: UnaryExample
Response: 
{
    "int32Field": 1,
    "basicTypes": {
       "intField": 2
    }
}
```
