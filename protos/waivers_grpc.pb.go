// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: waivers.proto

package protos

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

// WaiverClient is the client API for Waiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WaiverClient interface {
	GetPlayerStatus(ctx context.Context, in *GetPlayerStatusRequest, opts ...grpc.CallOption) (*GetPlayerStatusResponse, error)
}

type waiverClient struct {
	cc grpc.ClientConnInterface
}

func NewWaiverClient(cc grpc.ClientConnInterface) WaiverClient {
	return &waiverClient{cc}
}

func (c *waiverClient) GetPlayerStatus(ctx context.Context, in *GetPlayerStatusRequest, opts ...grpc.CallOption) (*GetPlayerStatusResponse, error) {
	out := new(GetPlayerStatusResponse)
	err := c.cc.Invoke(ctx, "/Waiver/GetPlayerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiverServer is the server API for Waiver service.
// All implementations must embed UnimplementedWaiverServer
// for forward compatibility
type WaiverServer interface {
	GetPlayerStatus(context.Context, *GetPlayerStatusRequest) (*GetPlayerStatusResponse, error)
	mustEmbedUnimplementedWaiverServer()
}

// UnimplementedWaiverServer must be embedded to have forward compatible implementations.
type UnimplementedWaiverServer struct {
}

func (UnimplementedWaiverServer) GetPlayerStatus(context.Context, *GetPlayerStatusRequest) (*GetPlayerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerStatus not implemented")
}
func (UnimplementedWaiverServer) mustEmbedUnimplementedWaiverServer() {}

// UnsafeWaiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WaiverServer will
// result in compilation errors.
type UnsafeWaiverServer interface {
	mustEmbedUnimplementedWaiverServer()
}

func RegisterWaiverServer(s grpc.ServiceRegistrar, srv WaiverServer) {
	s.RegisterService(&Waiver_ServiceDesc, srv)
}

func _Waiver_GetPlayerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiverServer).GetPlayerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Waiver/GetPlayerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiverServer).GetPlayerStatus(ctx, req.(*GetPlayerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Waiver_ServiceDesc is the grpc.ServiceDesc for Waiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Waiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Waiver",
	HandlerType: (*WaiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerStatus",
			Handler:    _Waiver_GetPlayerStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "waivers.proto",
}