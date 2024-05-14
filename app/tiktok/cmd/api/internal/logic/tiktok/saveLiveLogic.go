package tiktok

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/app/tiktok/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLiveLogic {
	return &SaveLiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLiveLogic) SaveLive(req *types.SaveLiveReq) (*types.SaveLiveResp, error) {
	result, err := l.svcCtx.TiktokRpc.SaveLive(l.ctx, &pb.SaveLiveReq{
		OpenUrl: req.OpenUrl,
		Name:    req.Name,
		Info:    req.Info,
	})
	if err != nil {
		return nil, err
	}
	redPacket, err := tranfRedPacketFromJson(req.Info)
	if err != nil {
		return nil, err
	}
	openTime := getOpenTime(redPacket)
	isNeedFork := checkNeedFork(redPacket)
	resp := &types.SaveLiveResp{}
	copier.Copy(resp, result)
	m := map[string]interface{}{
		"openTime":   openTime,
		"isNeedFork": isNeedFork,
	}
	jsonStr, _ := json.Marshal(m)
	resp.Content = string(jsonStr)
	return resp, nil
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
		if item == "加入粉丝团" || strings.Contains(item, "粉丝团") || item == "分享直播间" {
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
