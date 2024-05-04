package logic

import (
	"context"

	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLiveOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLiveOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLiveOneLogic {
	return &DeleteLiveOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLiveOneLogic) DeleteLiveOne(in *pb.DeleteLiveReq) (*pb.DeleteLiveResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.TiktokModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to delete tiktok live err : %v , in :%+v", err, in)
	}
	logx.Info("deleteLive for Id: ", in.Id)
	return &pb.DeleteLiveResp{
		Id: in.Id,
	}, nil
}
