package svc

import (
	"TikTokLite/app/message/cmd/api/internal/config"
	"TikTokLite/app/message/cmd/rpc/message"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	MessageRpcClient message.MessageZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		MessageRpcClient: message.NewMessageZrpcClient(zrpc.MustNewClient(c.MessageRpcConfig)),
	}
}
