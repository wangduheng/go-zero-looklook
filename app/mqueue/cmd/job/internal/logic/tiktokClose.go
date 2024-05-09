package logic

import (
	"context"
	"encoding/json"
	"io"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/tiktok/cmd/rpc/tiktok"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
)

// SettleRecordHandler   shcedule billing to home business
type TiktokRecordCloseHandler struct {
	svcCtx *svc.ServiceContext
}
type TiktokRecordOpenHandler struct {
	svcCtx *svc.ServiceContext
}

func NewTiktokRecordCloseHandler(svcCtx *svc.ServiceContext) *TiktokRecordCloseHandler {
	return &TiktokRecordCloseHandler{
		svcCtx: svcCtx,
	}
}

func NewTiktokRecordOpenHandler(svcCtx *svc.ServiceContext) *TiktokRecordOpenHandler {
	return &TiktokRecordOpenHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *TiktokRecordCloseHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.DeferCloseTiktokRecord
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "DeferCloseTiktokRecord payload err:%v, payLoad:%+v", err, t.Payload())
	}
	logx.Error("will close tiktok lastId:", p.ID)
	resp, err := l.svcCtx.TiktokRpc.DeleteLiveOne(ctx, &tiktok.DeleteLiveReq{
		Id: p.ID,
	})
	if err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "DeferCloseTiktokRecord  DeleteLiveOne fail or  no exists err:%v, sn:%d ,HomestayOrder : %+v", err, p.ID, resp)
	}

	return nil
}

// every one minute exec : if return err != nil , asynq will retry
func (l *TiktokRecordOpenHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.DeferOpenTiktokRecord
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "DeferOpenTiktokRecord payload err:%v, payLoad:%+v", err, t.Payload())
	}
	resp, err := l.svcCtx.TiktokRpc.FindLiveOne(ctx, &tiktok.FindLiveReq{
		Id: p.ID,
	})
	logx.Error("will open tiktok ", resp)

	if err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "TiktokRecordOpenHandler  FindLiveOne fail or  no exists err:%v, sn:%d ,HomestayOrder : %+v", err, p.ID, resp)
	}
	postData := struct {
		Id int64 `json:"Id"`
	}{
		Id: p.ID,
	}

	respHttpc, err := httpc.Do(context.Background(), http.MethodPost, "http://looklook:1005/tiktok/v1/receive", postData)

	if err != nil {
		logx.Error("ProcessTask httpc error", err)
	} else {
		bodyBytes, err := io.ReadAll(respHttpc.Body)
		if err != nil {
			// 处理读取响应体时的错误
			logx.Error("ProcessTask ReadAll error", err)
		} else {
			// 将响应体字节转换为字符串
			logx.Error("ProcessTask httpx ", string(bodyBytes))
		}

	}

	return nil
}
