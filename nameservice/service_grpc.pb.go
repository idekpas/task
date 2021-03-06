// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// NameServiceClient is the client API for NameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameServiceClient interface {
	GetName(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NameResponse, error)
	SetName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*Empty, error)
}

type nameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNameServiceClient(cc grpc.ClientConnInterface) NameServiceClient {
	return &nameServiceClient{cc}
}

func (c *nameServiceClient) GetName(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := c.cc.Invoke(ctx, "/NameService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameServiceClient) SetName(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/NameService/SetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameServiceServer is the server API for NameService service.
// All implementations must embed UnimplementedNameServiceServer
// for forward compatibility
type NameServiceServer interface {
	GetName(context.Context, *Empty) (*NameResponse, error)
	SetName(context.Context, *NameRequest) (*Empty, error)
	mustEmbedUnimplementedNameServiceServer()
}

// UnimplementedNameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNameServiceServer struct {
}

func (UnimplementedNameServiceServer) GetName(context.Context, *Empty) (*NameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedNameServiceServer) SetName(context.Context, *NameRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetName not implemented")
}
func (UnimplementedNameServiceServer) mustEmbedUnimplementedNameServiceServer() {}

// UnsafeNameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameServiceServer will
// result in compilation errors.
type UnsafeNameServiceServer interface {
	mustEmbedUnimplementedNameServiceServer()
}

func RegisterNameServiceServer(s grpc.ServiceRegistrar, srv NameServiceServer) {
	s.RegisterService(&NameService_ServiceDesc, srv)
}

func _NameService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).GetName(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameService_SetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).SetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameService/SetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).SetName(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NameService_ServiceDesc is the grpc.ServiceDesc for NameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NameService",
	HandlerType: (*NameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _NameService_GetName_Handler,
		},
		{
			MethodName: "SetName",
			Handler:    _NameService_SetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
