// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/pb/source.proto

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

// SourceClient is the client API for Source service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceClient interface {
	// Get all tables the source plugin supports
	GetTables(ctx context.Context, in *GetTables_Request, opts ...grpc.CallOption) (*GetTables_Response, error)
	// Get an example configuration for the source plugin
	GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error)
	// Configure the source plugin with the given spec
	Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error)
	// Fetch resources
	Fetch(ctx context.Context, in *Fetch_Request, opts ...grpc.CallOption) (Source_FetchClient, error)
}

type sourceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceClient(cc grpc.ClientConnInterface) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) GetTables(ctx context.Context, in *GetTables_Request, opts ...grpc.CallOption) (*GetTables_Response, error) {
	out := new(GetTables_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error) {
	out := new(GetExampleConfig_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/GetExampleConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error) {
	out := new(Configure_Response)
	err := c.cc.Invoke(ctx, "/proto.Source/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Fetch(ctx context.Context, in *Fetch_Request, opts ...grpc.CallOption) (Source_FetchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Source_ServiceDesc.Streams[0], "/proto.Source/Fetch", opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceFetchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Source_FetchClient interface {
	Recv() (*Fetch_Response, error)
	grpc.ClientStream
}

type sourceFetchClient struct {
	grpc.ClientStream
}

func (x *sourceFetchClient) Recv() (*Fetch_Response, error) {
	m := new(Fetch_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SourceServer is the server API for Source service.
// All implementations must embed UnimplementedSourceServer
// for forward compatibility
type SourceServer interface {
	// Get all tables the source plugin supports
	GetTables(context.Context, *GetTables_Request) (*GetTables_Response, error)
	// Get an example configuration for the source plugin
	GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error)
	// Configure the source plugin with the given spec
	Configure(context.Context, *Configure_Request) (*Configure_Response, error)
	// Fetch resources
	Fetch(*Fetch_Request, Source_FetchServer) error
	mustEmbedUnimplementedSourceServer()
}

// UnimplementedSourceServer must be embedded to have forward compatible implementations.
type UnimplementedSourceServer struct {
}

func (UnimplementedSourceServer) GetTables(context.Context, *GetTables_Request) (*GetTables_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTables not implemented")
}
func (UnimplementedSourceServer) GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleConfig not implemented")
}
func (UnimplementedSourceServer) Configure(context.Context, *Configure_Request) (*Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedSourceServer) Fetch(*Fetch_Request, Source_FetchServer) error {
	return status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedSourceServer) mustEmbedUnimplementedSourceServer() {}

// UnsafeSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceServer will
// result in compilation errors.
type UnsafeSourceServer interface {
	mustEmbedUnimplementedSourceServer()
}

func RegisterSourceServer(s grpc.ServiceRegistrar, srv SourceServer) {
	s.RegisterService(&Source_ServiceDesc, srv)
}

func _Source_GetTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTables_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetTables(ctx, req.(*GetTables_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_GetExampleConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExampleConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).GetExampleConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/GetExampleConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).GetExampleConfig(ctx, req.(*GetExampleConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Source/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Configure(ctx, req.(*Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Fetch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Fetch_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceServer).Fetch(m, &sourceFetchServer{stream})
}

type Source_FetchServer interface {
	Send(*Fetch_Response) error
	grpc.ServerStream
}

type sourceFetchServer struct {
	grpc.ServerStream
}

func (x *sourceFetchServer) Send(m *Fetch_Response) error {
	return x.ServerStream.SendMsg(m)
}

// Source_ServiceDesc is the grpc.ServiceDesc for Source service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Source_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Source",
	HandlerType: (*SourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTables",
			Handler:    _Source_GetTables_Handler,
		},
		{
			MethodName: "GetExampleConfig",
			Handler:    _Source_GetExampleConfig_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Source_Configure_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Fetch",
			Handler:       _Source_Fetch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pb/source.proto",
}
