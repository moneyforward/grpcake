package grpcake

import (
	"context"
	"errors"
	"fmt"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/linker"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// DescriptorSource is a source of protobuf descriptor information. It can be backed by a FileDescriptorSet
// proto (like a file generated by protoc) or a remote server that supports the reflection API.
type DescriptorSource interface {
	// FindSymbol returns a descriptor for the given fully-qualified symbol name.
	FindServiceDescriptor(fullyQualifiedName string) (protoreflect.ServiceDescriptor, error)
}

type fileSource struct {
	files linker.Files
}

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
	return fileSource{files: files}, nil
}

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

type serverSource struct {
	client *grpcreflect.Client
}

// DescriptorSourceFromServer creates a DescriptorSource that uses the given gRPC reflection client
// to interrogate a server for descriptor information. If the server does not support the reflection
// API then the various DescriptorSource methods will return ErrReflectionNotSupported
func DescriptorSourceFromServer(_ context.Context, refClient *grpcreflect.Client) DescriptorSource {
	return serverSource{client: refClient}
}

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
