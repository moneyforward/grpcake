# grpcake

A deliciously simple gRPC client written in go, `grpcake` allows you test your gRPC 
client from the command line without stressing about writing correct raw json strings.

## Usage 

```
Usage of bin/grpcake:
  -body string
        JSON body
  -grpc-method string
        GRPC Method
  -import string
        Proto files to import
  -url string
        GRPC Server URL
  -use-reflection
        Try to use server reflection (default true)
```

## Build

Run in a terminal

``` sh
make build
```
to bake your `grpcake` in _bin/grpcake_.

## Examples

Using the test service in this repository

```sh
make run-test-server REFLECT=false
```

Change the value of `REFLECT` to `true` to run test server with reflection.

### With reflection

```sh
$ bin/grpcake --url localhost:6069 -use-reflection=true\
--grpc-method testing.ExampleService/UnaryExample \
long_field:=10 int_field:=1 float_field:=1.5 double_field:=2.5 bool_field:=true \
string_field="hello world" bytes_field="b25lcGllY2VraW5kYXN1Y2sK"

2023/04/25 12:46:00 request json body: %s {"long_field":10,"int_field":1,"float_field":1.5,"double_field":2.5,"bool_field":true,"string_field":"hello world","bytes_field":"b25lcGllY2VraW5kYXN1Y2sK"}
2023/04/25 12:46:00 Response:  {
        "intField": 1,
        "longField": "10",
        "floatField": 1.5,
        "doubleField": 2.5,
        "boolField": true,
        "stringField": "hello world",
        "bytesField": "b25lcGllY2VraW5kYXN1Y2sK"
}
```

### Without reflection
```sh
$ bin/grpcake --url localhost:6069 -use-reflectin=false\
--grpc-method testing.ExampleService/UnaryExample \
--import internal/testing/example.proto \
long_field:=10 int_field:=1 float_field:=1.5 double_field:=2.5 bool_field:=true \
string_field="hello world" bytes_field="b25lcGllY2VraW5kYXN1Y2sK"

2023/04/25 12:46:00 request json body: %s {"long_field":10,"int_field":1,"float_field":1.5,"double_field":2.5,"bool_field":true,"string_field":"hello world","bytes_field":"b25lcGllY2VraW5kYXN1Y2sK"}
2023/04/25 12:46:00 Response:  {
        "intField": 1,
        "longField": "10",
        "floatField": 1.5,
        "doubleField": 2.5,
        "boolField": true,
        "stringField": "hello world",
        "bytesField": "b25lcGllY2VraW5kYXN1Y2sK"
}
```

## Contributing

Install go v1.20.3 (with [goenv](https://github.com/syndbg/goenv) recommended) then
install the necessary tools in _bin/_ with by running

```sh
make install-tools
```

### To run example server

Using the test service in this repository
```sh
make run-test-server REFLECT=false
```
Change the value of `REFLECT` to `true` to run test server with reflection.

### To regenerate `example.proto` files in `internal/testing` folder

 ```sh
 make gen-testing-proto
 ```
