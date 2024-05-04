package tiktok

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tiktok.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
