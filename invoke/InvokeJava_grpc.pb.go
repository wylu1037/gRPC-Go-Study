// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/InvokeJava.proto

package invoke

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

// IInvokeJavaClient is the client API for IInvokeJava service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IInvokeJavaClient interface {
	InvokeCrossLanguage(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
}

type iInvokeJavaClient struct {
	cc grpc.ClientConnInterface
}

func NewIInvokeJavaClient(cc grpc.ClientConnInterface) IInvokeJavaClient {
	return &iInvokeJavaClient{cc}
}

func (c *iInvokeJavaClient) InvokeCrossLanguage(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := c.cc.Invoke(ctx, "/invoke.IInvokeJava/invokeCrossLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IInvokeJavaServer is the server API for IInvokeJava service.
// All implementations must embed UnimplementedIInvokeJavaServer
// for forward compatibility
type IInvokeJavaServer interface {
	InvokeCrossLanguage(context.Context, *InvokeRequest) (*InvokeResponse, error)
	mustEmbedUnimplementedIInvokeJavaServer()
}

// UnimplementedIInvokeJavaServer must be embedded to have forward compatible implementations.
type UnimplementedIInvokeJavaServer struct {
}

func (UnimplementedIInvokeJavaServer) InvokeCrossLanguage(context.Context, *InvokeRequest) (*InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeCrossLanguage not implemented")
}
func (UnimplementedIInvokeJavaServer) mustEmbedUnimplementedIInvokeJavaServer() {}

// UnsafeIInvokeJavaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IInvokeJavaServer will
// result in compilation errors.
type UnsafeIInvokeJavaServer interface {
	mustEmbedUnimplementedIInvokeJavaServer()
}

func RegisterIInvokeJavaServer(s grpc.ServiceRegistrar, srv IInvokeJavaServer) {
	s.RegisterService(&IInvokeJava_ServiceDesc, srv)
}

func _IInvokeJava_InvokeCrossLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IInvokeJavaServer).InvokeCrossLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoke.IInvokeJava/invokeCrossLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IInvokeJavaServer).InvokeCrossLanguage(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IInvokeJava_ServiceDesc is the grpc.ServiceDesc for IInvokeJava service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IInvokeJava_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "invoke.IInvokeJava",
	HandlerType: (*IInvokeJavaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "invokeCrossLanguage",
			Handler:    _IInvokeJava_InvokeCrossLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/InvokeJava.proto",
}
