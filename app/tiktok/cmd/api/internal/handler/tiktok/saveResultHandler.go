package tiktok

import (
	"net/http"

	"looklook/app/tiktok/cmd/api/internal/logic/tiktok"
	"looklook/app/tiktok/cmd/api/internal/svc"
	"looklook/app/tiktok/cmd/api/internal/types"
	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveResultReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := tiktok.NewSaveResultLogic(r.Context(), svcCtx)
		resp, err := l.SaveResult(&req)
		result.HttpResult(r, w, resp, err)
	}
}
