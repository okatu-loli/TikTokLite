package router

import (
	"TikTokLite/api/internal/handler"
	"TikTokLite/api/internal/middleware"
	"net/http"

	"TikTokLite/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterAllHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 获取用户信息
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/user/",
				Handler: middleware.JwtMiddleware(handler.UserInfoHandler(serverCtx)),
			},
		},
	)
	// 登录
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login/",
				Handler: handler.LoginHandler(serverCtx),
			},
		},
	)
	// 注册
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register/",
				Handler: handler.RegisterHandler(serverCtx),
			},
		},
	)
}
