package tiktok

import (
	"net/http"

	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tiktok.NewSocketLogic(r.Context(), svcCtx)
		l.ServeWs(w, r)
	}
}

func ReceiveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReceiveReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := tiktok.NewSocketLogic(r.Context(), svcCtx)
		err := l.ResultsHandler(&req)
		result.HttpResult(r, w, "okok", err)
	}
}
