// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: crivo.proto

package gen

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

// CrivoClient is the client API for Crivo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrivoClient interface {
	Crivo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type crivoClient struct {
	cc grpc.ClientConnInterface
}

func NewCrivoClient(cc grpc.ClientConnInterface) CrivoClient {
	return &crivoClient{cc}
}

func (c *crivoClient) Crivo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/crivo.Crivo/crivo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrivoServer is the server API for Crivo service.
// All implementations must embed UnimplementedCrivoServer
// for forward compatibility
type CrivoServer interface {
	Crivo(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedCrivoServer()
}

// UnimplementedCrivoServer must be embedded to have forward compatible implementations.
type UnimplementedCrivoServer struct {
}

func (UnimplementedCrivoServer) Crivo(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Crivo not implemented")
}
func (UnimplementedCrivoServer) mustEmbedUnimplementedCrivoServer() {}

// UnsafeCrivoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrivoServer will
// result in compilation errors.
type UnsafeCrivoServer interface {
	mustEmbedUnimplementedCrivoServer()
}

func RegisterCrivoServer(s grpc.ServiceRegistrar, srv CrivoServer) {
	s.RegisterService(&Crivo_ServiceDesc, srv)
}

func _Crivo_Crivo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrivoServer).Crivo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crivo.Crivo/crivo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrivoServer).Crivo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Crivo_ServiceDesc is the grpc.ServiceDesc for Crivo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Crivo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crivo.Crivo",
	HandlerType: (*CrivoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "crivo",
			Handler:    _Crivo_Crivo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crivo.proto",
}
