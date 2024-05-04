package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/tiktok/cmd/rpc/tiktok"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

// SettleRecordHandler   shcedule billing to home business
type TiktokRecordHandler struct {
	svcCtx *svc.ServiceContext
}

func NewTiktokRecordHandler(svcCtx *svc.ServiceContext) *TiktokRecordHandler {
	return &TiktokRecordHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *TiktokRecordHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	fmt.Printf("shcedule job delete tiktok \n")

	var p jobtype.DeferCloseTiktokRecord
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "DeferCloseTiktokRecord payload err:%v, payLoad:%+v", err, t.Payload())
	}

	resp, err := l.svcCtx.TiktokRpc.DeleteLiveOne(ctx, &tiktok.DeleteLiveReq{
		Id: p.ID,
	})
	if err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "DeferCloseTiktokRecord  DeleteLiveOne fail or  no exists err:%v, sn:%d ,HomestayOrder : %+v", err, p.ID, resp)
	}

	return nil
}
