package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	SecretKey      string `yaml:"jwt_secret_key"`
	DataSourceName string
}
