package tiktok

import (
	"context"

	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FetchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchLogic {
	return &FetchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchLogic) Fetch() (*types.FetchLiveResp, error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	logx.Error("userId:", userId)
	result, err := l.svcCtx.TiktokRpc.FetchLives(l.ctx, &pb.FetchLiveReq{
		UserId:     userId,
		LastSecond: 60,
	})
	if err != nil {
		return nil, err
	}
	var lives []types.TitokLive
	if len(result.List) > 0 {
		for _, item := range result.List {
			var live types.TitokLive
			_ = copier.Copy(&live, item)
			lives = append(lives, live)
		}
	}
	return &types.FetchLiveResp{List: lives}, nil
}
