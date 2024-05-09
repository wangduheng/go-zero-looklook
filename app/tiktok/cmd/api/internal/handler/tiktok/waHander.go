package tiktok

import (
	"net/http"

	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
			logx.Error("Parse err: ", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := tiktok.NewSocketLogic(r.Context(), svcCtx)
		err := l.ResultsHandler(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, "okok")
		}
	}
}
