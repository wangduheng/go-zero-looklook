package tiktok

import (
	"context"

	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/app/tiktok/cmd/rpc/tiktok"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveResultLogic {
	return &SaveResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveResultLogic) SaveResult(req *types.SaveResultReq) (resp *types.SaveResultResp, err error) {
	rpcResp, err := l.svcCtx.TiktokRpc.SaveResult(l.ctx, &tiktok.SaveResultReq{
		Id:   req.Id,
		Name: req.Name,
	})
	if err == nil {
		resp = &types.SaveResultResp{
			Content: rpcResp.Content,
		}
	}
	logx.Error("SaveResult: ", resp.Content)
	return
}
