# grpcake

```sh
➜  grpcake git:(main) ✗ go run cmd/main.go --url localhost:6069 \
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

```sh
make run-test-server
```

### To regenerate `example.proto` files in `internal/testing` folder

 ```sh
 bin/buf generate --template testing/buf.gen.yaml;
 ```
