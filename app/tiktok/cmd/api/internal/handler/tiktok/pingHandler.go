package tiktok

import (
	"net/http"

	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/common/result"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tiktok.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		result.HttpResult(r, w, resp, err)
	}
}
