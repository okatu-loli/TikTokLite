package svc

import (
	"TikTokLite/app/message/cmd/rpc/internal/config"
	"TikTokLite/app/message/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"log"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(db, c.Cache),
	}
}
