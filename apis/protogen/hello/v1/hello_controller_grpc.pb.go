// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: hello/v1/hello_controller.proto

package hellov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HelloController_SayHello_FullMethodName = "/hello.v1.HelloController/SayHello"
)

// HelloControllerClient is the client API for HelloController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloControllerClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type helloControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloControllerClient(cc grpc.ClientConnInterface) HelloControllerClient {
	return &helloControllerClient{cc}
}

func (c *helloControllerClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, HelloController_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloControllerServer is the server API for HelloController service.
// All implementations must embed UnimplementedHelloControllerServer
// for forward compatibility.
type HelloControllerServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedHelloControllerServer()
}

// UnimplementedHelloControllerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHelloControllerServer struct{}

func (UnimplementedHelloControllerServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloControllerServer) mustEmbedUnimplementedHelloControllerServer() {}
func (UnimplementedHelloControllerServer) testEmbeddedByValue()                         {}

// UnsafeHelloControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloControllerServer will
// result in compilation errors.
type UnsafeHelloControllerServer interface {
	mustEmbedUnimplementedHelloControllerServer()
}

func RegisterHelloControllerServer(s grpc.ServiceRegistrar, srv HelloControllerServer) {
	// If the following call pancis, it indicates UnimplementedHelloControllerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HelloController_ServiceDesc, srv)
}

func _HelloController_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloControllerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloController_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloControllerServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloController_ServiceDesc is the grpc.ServiceDesc for HelloController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.v1.HelloController",
	HandlerType: (*HelloControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloController_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/v1/hello_controller.proto",
}
