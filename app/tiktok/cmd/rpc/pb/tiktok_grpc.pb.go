// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: tiktok.proto

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

const (
	Tiktok_SaveLive_FullMethodName      = "/pb.tiktok/saveLive"
	Tiktok_FetchLives_FullMethodName    = "/pb.tiktok/fetchLives"
	Tiktok_DeleteLiveOne_FullMethodName = "/pb.tiktok/DeleteLiveOne"
	Tiktok_FindLiveOne_FullMethodName   = "/pb.tiktok/FindLiveOne"
)

// TiktokClient is the client API for Tiktok service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TiktokClient interface {
	SaveLive(ctx context.Context, in *SaveLiveReq, opts ...grpc.CallOption) (*SaveLiveResp, error)
	FetchLives(ctx context.Context, in *FetchLiveReq, opts ...grpc.CallOption) (*FetchLiveResp, error)
	DeleteLiveOne(ctx context.Context, in *DeleteLiveReq, opts ...grpc.CallOption) (*DeleteLiveResp, error)
	FindLiveOne(ctx context.Context, in *FindLiveReq, opts ...grpc.CallOption) (*FindLiveResp, error)
}

type tiktokClient struct {
	cc grpc.ClientConnInterface
}

func NewTiktokClient(cc grpc.ClientConnInterface) TiktokClient {
	return &tiktokClient{cc}
}

func (c *tiktokClient) SaveLive(ctx context.Context, in *SaveLiveReq, opts ...grpc.CallOption) (*SaveLiveResp, error) {
	out := new(SaveLiveResp)
	err := c.cc.Invoke(ctx, Tiktok_SaveLive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiktokClient) FetchLives(ctx context.Context, in *FetchLiveReq, opts ...grpc.CallOption) (*FetchLiveResp, error) {
	out := new(FetchLiveResp)
	err := c.cc.Invoke(ctx, Tiktok_FetchLives_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiktokClient) DeleteLiveOne(ctx context.Context, in *DeleteLiveReq, opts ...grpc.CallOption) (*DeleteLiveResp, error) {
	out := new(DeleteLiveResp)
	err := c.cc.Invoke(ctx, Tiktok_DeleteLiveOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tiktokClient) FindLiveOne(ctx context.Context, in *FindLiveReq, opts ...grpc.CallOption) (*FindLiveResp, error) {
	out := new(FindLiveResp)
	err := c.cc.Invoke(ctx, Tiktok_FindLiveOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TiktokServer is the server API for Tiktok service.
// All implementations must embed UnimplementedTiktokServer
// for forward compatibility
type TiktokServer interface {
	SaveLive(context.Context, *SaveLiveReq) (*SaveLiveResp, error)
	FetchLives(context.Context, *FetchLiveReq) (*FetchLiveResp, error)
	DeleteLiveOne(context.Context, *DeleteLiveReq) (*DeleteLiveResp, error)
	FindLiveOne(context.Context, *FindLiveReq) (*FindLiveResp, error)
	mustEmbedUnimplementedTiktokServer()
}

// UnimplementedTiktokServer must be embedded to have forward compatible implementations.
type UnimplementedTiktokServer struct {
}

func (UnimplementedTiktokServer) SaveLive(context.Context, *SaveLiveReq) (*SaveLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLive not implemented")
}
func (UnimplementedTiktokServer) FetchLives(context.Context, *FetchLiveReq) (*FetchLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLives not implemented")
}
func (UnimplementedTiktokServer) DeleteLiveOne(context.Context, *DeleteLiveReq) (*DeleteLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLiveOne not implemented")
}
func (UnimplementedTiktokServer) FindLiveOne(context.Context, *FindLiveReq) (*FindLiveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLiveOne not implemented")
}
func (UnimplementedTiktokServer) mustEmbedUnimplementedTiktokServer() {}

// UnsafeTiktokServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TiktokServer will
// result in compilation errors.
type UnsafeTiktokServer interface {
	mustEmbedUnimplementedTiktokServer()
}

func RegisterTiktokServer(s grpc.ServiceRegistrar, srv TiktokServer) {
	s.RegisterService(&Tiktok_ServiceDesc, srv)
}

func _Tiktok_SaveLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiktokServer).SaveLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiktok_SaveLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiktokServer).SaveLive(ctx, req.(*SaveLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiktok_FetchLives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiktokServer).FetchLives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiktok_FetchLives_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiktokServer).FetchLives(ctx, req.(*FetchLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiktok_DeleteLiveOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiktokServer).DeleteLiveOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiktok_DeleteLiveOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiktokServer).DeleteLiveOne(ctx, req.(*DeleteLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tiktok_FindLiveOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TiktokServer).FindLiveOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tiktok_FindLiveOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TiktokServer).FindLiveOne(ctx, req.(*FindLiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tiktok_ServiceDesc is the grpc.ServiceDesc for Tiktok service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tiktok_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.tiktok",
	HandlerType: (*TiktokServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "saveLive",
			Handler:    _Tiktok_SaveLive_Handler,
		},
		{
			MethodName: "fetchLives",
			Handler:    _Tiktok_FetchLives_Handler,
		},
		{
			MethodName: "DeleteLiveOne",
			Handler:    _Tiktok_DeleteLiveOne_Handler,
		},
		{
			MethodName: "FindLiveOne",
			Handler:    _Tiktok_FindLiveOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tiktok.proto",
}
