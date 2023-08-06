package svc

import (
	"TikTokLite/app/message/cmd/api/internal/config"
	"TikTokLite/app/message/cmd/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	MessageMiddleWare rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		MessageMiddleWare: middleware.NewMessageMiddleWareMiddleware().Handle,
	}
}
