package svc

import (
	"TikTokLite/api/internal/config"
	"TikTokLite/api/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ServiceContext struct {
	Config    config.Config
	DbEngin   *gorm.DB
	SecretKey string
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 添加Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "duck_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构
	autoMigrateErr := db.AutoMigrate(&model.AuthUser{}, &model.UserInfo{})
	if autoMigrateErr != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		DbEngin:   db,
		SecretKey: c.SecretKey,
	}
}
