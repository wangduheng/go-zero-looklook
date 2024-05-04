package logic

import (
	"context"
	"fmt"
	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/app/tiktok/model"
	"looklook/common/xerr"
	"strconv"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchLivesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFetchLivesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchLivesLogic {
	return &FetchLivesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *FetchLivesLogic) getLastFetchTime(userId int64) time.Time {
	currentTime := time.Now()

	lastFetchTimeStr, err := l.svcCtx.RedisClient.GetSet(fmt.Sprintf("LastFetchTime_%d", userId), strconv.FormatInt(currentTime.Unix(), 10))
	if err != nil {
		logx.Error("获取redis时间错误:", err)
		lastFetchTimeStr = strconv.FormatInt(currentTime.Add(-10*time.Minute).Unix(), 10)
	}
	logx.Error("lastFetchTime:", lastFetchTimeStr)

	lastTimestamp, err := strconv.ParseInt(lastFetchTimeStr, 10, 64)
	if err != nil {
		logx.Error("time.Parse错误:", err)
		lastTimestamp = currentTime.Add(-10 * time.Minute).Unix()
	}
	lastFetchTime := time.Unix(lastTimestamp, 0)
	return lastFetchTime
}
func (l *FetchLivesLogic) FetchLives(in *pb.FetchLiveReq) (*pb.FetchLiveResp, error) {
	currentTime := time.Now()
	lastFetchTime := l.getLastFetchTime(in.UserId)
	timeToCompare := currentTime.Add(time.Duration(in.LastSecond) * time.Second)
	whereBuilder := l.svcCtx.TiktokModel.SelectBuilder().Where(squirrel.Gt{"open_time": currentTime}).Where(squirrel.Lt{"open_time": timeToCompare})
	whereBuilder = whereBuilder.Where(squirrel.Gt{"create_time": lastFetchTime})
	list, err := l.svcCtx.TiktokModel.FindAll(l.ctx, whereBuilder, "open_time")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}
	var resp []*pb.LiveData
	if len(list) > 0 {
		for _, item := range list {
			var livedata pb.LiveData
			_ = copier.Copy(&livedata, item)
			livedata.CreateTime = item.CreateTime.Unix()
			livedata.OpenTime = item.OpenTime.Unix()
			livedata.IsNeedFork = item.IsNeedFork
			resp = append(resp, &livedata)
		}
	}
	return &pb.FetchLiveResp{List: resp}, nil
}
