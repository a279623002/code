package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// DBConfig 数据库配置
type DBConfig struct {
	DataSource      string `json:"DataSource"`
	MaxIdleConns    int    `json:"MaxIdleConns,default=10"`
	MaxOpenConns    int    `json:"MaxOpenConns,default=100"`
	ConnMaxLifetime int    `json:"ConnMaxLifetime,default=3600"`
}

// Config user api 配置
type Config struct {
	rest.RestConf
	DB       DBConfig
	OrderRpc zrpc.RpcClientConf
}
