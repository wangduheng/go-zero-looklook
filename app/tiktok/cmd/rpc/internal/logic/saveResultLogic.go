package logic

import (
	"context"

	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/app/tiktok/model"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveResultLogic {
	return &SaveResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveResultLogic) SaveResult(in *pb.SaveResultReq) (*pb.SaveResultResp, error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.LiveResultModel.Insert(l.ctx, &model.LiveResult{
		Uid:  userId,
		Lid:  in.Id,
		Name: in.Name,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SaveResult db  err:%v,user:%+v", err, in)
	}
	return &pb.SaveResultResp{
		Content: "ok",
	}, nil
}
