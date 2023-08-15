package handler

import (
	"net/http"

	"douyinFavoriteList_4/api/internal/logic"
	"douyinFavoriteList_4/api/internal/svc"
	"douyinFavoriteList_4/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.GetFavoriteList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
