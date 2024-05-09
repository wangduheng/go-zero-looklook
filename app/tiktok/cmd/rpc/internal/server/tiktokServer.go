// Code generated by goctl. DO NOT EDIT.
// Source: tiktok.proto

package server

import (
	"context"

	"looklook/app/tiktok/cmd/rpc/internal/logic"
	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/pb"
)

type TiktokServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTiktokServer
}

func NewTiktokServer(svcCtx *svc.ServiceContext) *TiktokServer {
	return &TiktokServer{
		svcCtx: svcCtx,
	}
}

func (s *TiktokServer) SaveLive(ctx context.Context, in *pb.SaveLiveReq) (*pb.SaveLiveResp, error) {
	l := logic.NewSaveLiveLogic(ctx, s.svcCtx)
	return l.SaveLive(in)
}

func (s *TiktokServer) FetchLives(ctx context.Context, in *pb.FetchLiveReq) (*pb.FetchLiveResp, error) {
	l := logic.NewFetchLivesLogic(ctx, s.svcCtx)
	return l.FetchLives(in)
}

func (s *TiktokServer) DeleteLiveOne(ctx context.Context, in *pb.DeleteLiveReq) (*pb.DeleteLiveResp, error) {
	l := logic.NewDeleteLiveOneLogic(ctx, s.svcCtx)
	return l.DeleteLiveOne(in)
}

func (s *TiktokServer) FindLiveOne(ctx context.Context, in *pb.FindLiveReq) (*pb.FindLiveResp, error) {
	l := logic.NewFindLiveOneLogic(ctx, s.svcCtx)
	return l.FindLiveOne(in)
}