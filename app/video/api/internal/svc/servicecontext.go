package svc

import (
	"tiktoklite/app/video/api/internal/config"
	"tiktoklite/app/video/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//VideoRpc   videoservice.Videoservice
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlconn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		//VideoRpc:   videoservice.NewVideoservice(zrpc.MustNewClient(c.VideoRpcConf)),
		VideoModel: model.NewVideoModel(sqlconn, c.Cache),
	}
}
