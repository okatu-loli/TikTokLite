package handler

import (
	"net/http"

	"tiktoklite/app/video/api/internal/logic"
	"tiktoklite/app/video/api/internal/svc"
	"tiktoklite/app/video/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func publishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		//往logic里传递req
		l := logic.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req, file, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
