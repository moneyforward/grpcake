package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/grpc"
)

func main() {
	parser := protoparse.Parser{}
	fileDescriptors, err := parser.ParseFiles("cmd/adder.proto")
	if err != nil {
		log.Fatal(err)
	}

	fileDesc := fileDescriptors[0]

	if err != nil {
		log.Fatalf("error loading proto file descriptor: %v", err)
	}

	msgDesc := fileDesc.FindMessage("api.AddRequest")
	serviceDesc := fileDesc.FindService("api.Adder")

	reqMsg := dynamic.NewMessage(msgDesc)

	reqMsg.SetField(msgDesc.FindFieldByName("x"), int32(1))
	reqMsg.SetField(msgDesc.FindFieldByName("y"), int32(2))

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to grpc server: %v", err)
	}

	addMethodDesc := serviceDesc.FindMethodByName("Add")

	stubServer := grpcdynamic.NewStub(conn)
	resMsg, err := stubServer.InvokeRpc(context.Background(), addMethodDesc, reqMsg)
	if err != nil {
		log.Fatalf("error invoking grpc method: %v", err)
	}

	marshaler := jsonpb.Marshaler{}
	log.Println(marshaler.MarshalToString(resMsg))

}
