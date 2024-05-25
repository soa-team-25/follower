// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc3
// source: proto/follower/follower.proto

package follower

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

const (
	Follower_CreateNewFollowing_FullMethodName     = "/Follower/CreateNewFollowing"
	Follower_GetUserRecommendations_FullMethodName = "/Follower/GetUserRecommendations"
	Follower_FindUserFollowings_FullMethodName     = "/Follower/FindUserFollowings"
)

// FollowerClient is the client API for Follower service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowerClient interface {
	CreateNewFollowing(ctx context.Context, in *UserFollowingDto, opts ...grpc.CallOption) (*NeoFollowerDto, error)
	GetUserRecommendations(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListNeoUserDto, error)
	FindUserFollowings(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListNeoUserDto, error)
}

type followerClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowerClient(cc grpc.ClientConnInterface) FollowerClient {
	return &followerClient{cc}
}

func (c *followerClient) CreateNewFollowing(ctx context.Context, in *UserFollowingDto, opts ...grpc.CallOption) (*NeoFollowerDto, error) {
	out := new(NeoFollowerDto)
	err := c.cc.Invoke(ctx, Follower_CreateNewFollowing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerClient) GetUserRecommendations(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListNeoUserDto, error) {
	out := new(ListNeoUserDto)
	err := c.cc.Invoke(ctx, Follower_GetUserRecommendations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followerClient) FindUserFollowings(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ListNeoUserDto, error) {
	out := new(ListNeoUserDto)
	err := c.cc.Invoke(ctx, Follower_FindUserFollowings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowerServer is the server API for Follower service.
// All implementations must embed UnimplementedFollowerServer
// for forward compatibility
type FollowerServer interface {
	CreateNewFollowing(context.Context, *UserFollowingDto) (*NeoFollowerDto, error)
	GetUserRecommendations(context.Context, *Id) (*ListNeoUserDto, error)
	FindUserFollowings(context.Context, *Id) (*ListNeoUserDto, error)
	mustEmbedUnimplementedFollowerServer()
}

// UnimplementedFollowerServer must be embedded to have forward compatible implementations.
type UnimplementedFollowerServer struct {
}

func (UnimplementedFollowerServer) CreateNewFollowing(context.Context, *UserFollowingDto) (*NeoFollowerDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewFollowing not implemented")
}
func (UnimplementedFollowerServer) GetUserRecommendations(context.Context, *Id) (*ListNeoUserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecommendations not implemented")
}
func (UnimplementedFollowerServer) FindUserFollowings(context.Context, *Id) (*ListNeoUserDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserFollowings not implemented")
}
func (UnimplementedFollowerServer) mustEmbedUnimplementedFollowerServer() {}

// UnsafeFollowerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowerServer will
// result in compilation errors.
type UnsafeFollowerServer interface {
	mustEmbedUnimplementedFollowerServer()
}

func RegisterFollowerServer(s grpc.ServiceRegistrar, srv FollowerServer) {
	s.RegisterService(&Follower_ServiceDesc, srv)
}

func _Follower_CreateNewFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFollowingDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerServer).CreateNewFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follower_CreateNewFollowing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerServer).CreateNewFollowing(ctx, req.(*UserFollowingDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follower_GetUserRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerServer).GetUserRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follower_GetUserRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerServer).GetUserRecommendations(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follower_FindUserFollowings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowerServer).FindUserFollowings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follower_FindUserFollowings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowerServer).FindUserFollowings(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// Follower_ServiceDesc is the grpc.ServiceDesc for Follower service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Follower_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Follower",
	HandlerType: (*FollowerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewFollowing",
			Handler:    _Follower_CreateNewFollowing_Handler,
		},
		{
			MethodName: "GetUserRecommendations",
			Handler:    _Follower_GetUserRecommendations_Handler,
		},
		{
			MethodName: "FindUserFollowings",
			Handler:    _Follower_FindUserFollowings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/follower/follower.proto",
}
