package handler

import (
	"TikTokLite/api/internal/logic"
	"TikTokLite/api/internal/types"
	"github.com/go-playground/validator/v10"
	"net/http"

	"TikTokLite/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		req.Username = r.URL.Query().Get("username")
		req.Password = r.URL.Query().Get("password")

		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
