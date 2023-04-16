# grpcake

```sh
➜  grpcake git:(main) ✗ go run cmd/main.go --url localhost:8080 --grpc-method api.Adder/Add --import cmd/adder.proto x:=1 y:=2

2023/04/16 22:15:20 request json body: %s {"x":1,"y":2}
2023/04/16 22:15:20 {"result":3}
```

```sh
➜  grpcake git:(main) ✗ go run cmd/main.go --url localhost:8080 --grpc-method api.Adder/Add --import cmd/adder.proto --body '{"x": 1, "y": 2}'                         
2023/04/16 22:21:28 request json body: %s {"x": 1, "y": 2}
2023/04/16 22:21:28 {"result":3}
```
