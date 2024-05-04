package tiktok

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
)

func SaveLiveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveLiveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tiktok.NewSaveLiveLogic(r.Context(), svcCtx)
		resp, err := l.SaveLive(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
