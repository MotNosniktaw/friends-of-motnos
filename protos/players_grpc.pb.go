// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: players.proto

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

// PlayerClient is the client API for Player service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerClient interface {
	Get(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error)
}

type playerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerClient(cc grpc.ClientConnInterface) PlayerClient {
	return &playerClient{cc}
}

func (c *playerClient) Get(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error) {
	out := new(GetPlayerResponse)
	err := c.cc.Invoke(ctx, "/Player/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServer is the server API for Player service.
// All implementations must embed UnimplementedPlayerServer
// for forward compatibility
type PlayerServer interface {
	Get(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error)
	mustEmbedUnimplementedPlayerServer()
}

// UnimplementedPlayerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerServer struct {
}

func (UnimplementedPlayerServer) Get(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPlayerServer) mustEmbedUnimplementedPlayerServer() {}

// UnsafePlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServer will
// result in compilation errors.
type UnsafePlayerServer interface {
	mustEmbedUnimplementedPlayerServer()
}

func RegisterPlayerServer(s grpc.ServiceRegistrar, srv PlayerServer) {
	s.RegisterService(&Player_ServiceDesc, srv)
}

func _Player_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Player/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServer).Get(ctx, req.(*GetPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Player_ServiceDesc is the grpc.ServiceDesc for Player service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Player_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Player",
	HandlerType: (*PlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Player_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "players.proto",
}