package logic

import (
	"context"

	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/app/tiktok/model"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindLiveOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLiveOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLiveOneLogic {
	return &FindLiveOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLiveOneLogic) FindLiveOne(in *pb.FindLiveReq) (*pb.FindLiveResp, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.TiktokModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get tiktok's  err : %v , in :%+v", err, in)
	}
	logx.Error("FindLiveOne: ", data)
	var resp pb.LiveData
	if data != nil {
		_ = copier.Copy(&resp, data)
		resp.OpenTime = data.OpenTime.Unix()
		resp.IsNeedFork = data.IsNeedFork

	}
	return &pb.FindLiveResp{LiveDate: &resp}, nil
}
