package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"looklook/app/mqueue/cmd/job/jobtype"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"looklook/app/tiktok/cmd/rpc/internal/svc"
	"looklook/app/tiktok/cmd/rpc/internal/types"
	"looklook/app/tiktok/cmd/rpc/pb"
	"looklook/app/tiktok/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const CloseTiktokTimeMinutes = 10 //defer close order time

func NewSaveLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLiveLogic {
	return &SaveLiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveLiveLogic) SaveLive(in *pb.SaveLiveReq) (*pb.SaveLiveResp, error) {
	url := tranfOpenUrl(in.OpenUrl)
	redPacket, err := tranfRedPacketFromJson(in.Info)
	if err != nil {
		return nil, err
	}
	openTime := getOpenTime(redPacket)
	isNeedFork := checkNeedFork(redPacket)
	insertResult, err := l.svcCtx.TiktokModel.Insert(l.ctx, &model.DouyinLive{
		OpenUrl:    url,
		Name:       in.Name,
		Info:       in.Info,
		OpenTime:   openTime,
		IsNeedFork: isNeedFork,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SaveLive db  err:%v,user:%+v", err, in)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SaveLive db  insertResult.LastInsertId err:%v,user:%+v", err, in)
	}
	//2、Delayed closing of order tasks.
	payload, err := json.Marshal(jobtype.DeferCloseTiktokRecord{ID: lastId})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create defer close tikok task json Marshal fail err :%+v , sn : %s", err, lastId)
	} else {
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.ScheduleTiktokCloseRecord, payload), asynq.ProcessIn(CloseTiktokTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("create defer close tikok task insert queue fail err :%+v , sn : %s", err, lastId)
		}
	}
	return &pb.SaveLiveResp{
		Id: lastId,
	}, nil
}

func tranfRedPacketFromJson(jsonData string) (types.RedPacket, error) {
	var redPacket types.RedPacket
	// 解析 JSON 数据到结构体
	err := json.Unmarshal([]byte(jsonData), &redPacket)
	return redPacket, err
}

func tranfOpenUrl(data string) string {
	openUrl := data
	urlPattern := `https?://[^\s]+`
	re := regexp.MustCompile(urlPattern)
	open := re.FindString(openUrl)
	return open
}

func getOpenTime(redPacket types.RedPacket) time.Time {
	minutes, second, _ := parseCountdownString(redPacket.Time)
	duration := time.Duration(minutes)*time.Minute + time.Duration(second)*time.Second
	currentTime := time.Now()
	resultTime := currentTime.Add(duration)
	return resultTime

}

func checkNeedFork(edPacket types.RedPacket) int64 {
	var res int64 = 0
	for _, item := range edPacket.JoinCondition {
		if item == "加入粉丝团" {
			res = 1
			break
		}
	}
	return res
}

func parseCountdownString(countdownStr string) (int, int, error) {
	var minutes, second int
	// 使用 fmt.Scanf 解析字符串
	_, err := fmt.Sscanf(countdownStr, "%d:%d", &minutes, &second)
	if err != nil {
		return 0, 0, err
	}
	return minutes, second, nil
}
