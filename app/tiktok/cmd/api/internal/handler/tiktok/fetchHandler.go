package tiktok

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
)

func FetchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tiktok.NewFetchLogic(r.Context(), svcCtx)
		resp, err := l.Fetch()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
