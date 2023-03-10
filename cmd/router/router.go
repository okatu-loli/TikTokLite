// Code generated by hertz generator.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/okatu-loli/TikTokLite/internal/handler"
	"github.com/okatu-loli/TikTokLite/internal/middleware"
)

// CustomizedRegister registers customize routers.
func CustomizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	douyin := r.Group("/douyin")
	{
		//使用中间件
		//douyin.Use(basic_auth.BasicAuth(map[string]string{"test": "test"}))
		user := douyin.Group("/user")
		{
			user.POST("/register/", handler.Register)
			//user.POST("/login", handler.Login)
			//user.GET("/", handler.GetUserInfo)
			user.POST("/login/", middleware.JwtMiddleware.LoginHandler)
			user.GET("", middleware.JwtMiddleware.MiddlewareFunc(), handler.GetUserInfo)
		}

		video := douyin.Group("/publish")
		{
			video.Use(middleware.JwtMiddleware.MiddlewareFunc())
			video.POST("/action/", handler.UploadVideo)
			video.GET("/list", handler.PublishList)
		}

		douyin.GET("/feed", handler.FeedList)

		relation := douyin.Group("/relation")
		{
			relation.Use(middleware.JwtMiddleware.MiddlewareFunc())
			relation.POST("/action/", handler.PostFollowActionHandler)

			relation.GET("/follow/list", handler.QueryFollowListHandler)

			relation.GET("/follower/list", handler.QueryFollowerHandler)

			relation.GET("/friend/list")
		}
	}
}
