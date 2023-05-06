package grpcdynamic

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Stub struct {
	channel grpc.ClientConnInterface
}

func NewStub(channel grpc.ClientConnInterface) Stub {
	return Stub{channel: channel}
}

// InvokeRpc calls unary RPC methods on the server.
func (s Stub) InvokeRpc(ctx context.Context, method protoreflect.MethodDescriptor, request proto.Message, opts ...grpc.CallOption) (proto.Message, error) {
	if method.IsStreamingClient() || method.IsStreamingServer() {
		return nil, fmt.Errorf("InvokeRpc is for unary methods; %q is %s", method.FullName(), methodType(method))
	}

	if err := checkMessageType(method.Input(), request); err != nil {
		return nil, fmt.Errorf("error checking message type: %v", err)
	}

	resp := dynamicpb.NewMessage(method.Output())
	if err := s.channel.Invoke(ctx, requestMethod(method), request, resp, opts...); err != nil {
		return nil, fmt.Errorf("error invoking rpc method: %v", err)
	}

	return resp, nil
}

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

func requestMethod(md protoreflect.MethodDescriptor) string {
	return fmt.Sprintf("/%s/%s", md.Parent().FullName(), md.Name())
}

func methodType(md protoreflect.MethodDescriptor) string {
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
