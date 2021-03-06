// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: playground.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlaygroundClient is the client API for Playground service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaygroundClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (Playground_SayHelloStreamClient, error)
}

type playgroundClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaygroundClient(cc grpc.ClientConnInterface) PlaygroundClient {
	return &playgroundClient{cc}
}

func (c *playgroundClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/Playground/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (Playground_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Playground_ServiceDesc.Streams[0], "/Playground/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &playgroundSayHelloStreamClient{stream}
	return x, nil
}

type Playground_SayHelloStreamClient interface {
	Send(*SayHelloRequest) error
	Recv() (*SayHelloResponse, error)
	grpc.ClientStream
}

type playgroundSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *playgroundSayHelloStreamClient) Send(m *SayHelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *playgroundSayHelloStreamClient) Recv() (*SayHelloResponse, error) {
	m := new(SayHelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlaygroundServer is the server API for Playground service.
// All implementations must embed UnimplementedPlaygroundServer
// for forward compatibility
type PlaygroundServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	SayHelloStream(Playground_SayHelloStreamServer) error
	mustEmbedUnimplementedPlaygroundServer()
}

// UnimplementedPlaygroundServer must be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServer struct {
}

func (UnimplementedPlaygroundServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedPlaygroundServer) SayHelloStream(Playground_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedPlaygroundServer) mustEmbedUnimplementedPlaygroundServer() {}

// UnsafePlaygroundServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaygroundServer will
// result in compilation errors.
type UnsafePlaygroundServer interface {
	mustEmbedUnimplementedPlaygroundServer()
}

func RegisterPlaygroundServer(s grpc.ServiceRegistrar, srv PlaygroundServer) {
	s.RegisterService(&Playground_ServiceDesc, srv)
}

func _Playground_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Playground/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlaygroundServer).SayHelloStream(&playgroundSayHelloStreamServer{stream})
}

type Playground_SayHelloStreamServer interface {
	Send(*SayHelloResponse) error
	Recv() (*SayHelloRequest, error)
	grpc.ServerStream
}

type playgroundSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *playgroundSayHelloStreamServer) Send(m *SayHelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *playgroundSayHelloStreamServer) Recv() (*SayHelloRequest, error) {
	m := new(SayHelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Playground_ServiceDesc is the grpc.ServiceDesc for Playground service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playground_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Playground",
	HandlerType: (*PlaygroundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Playground_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _Playground_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "playground.proto",
}
