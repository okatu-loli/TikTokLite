package handler

import (
	"TikTokLite/api/internal/logic"
	"TikTokLite/api/internal/types"
	"net/http"
	"strconv"

	"TikTokLite/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		userIdStr := r.URL.Query().Get("user_id")
		req.UserId, _ = strconv.Atoi(userIdStr)
		req.Token = r.URL.Query().Get("token")

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
