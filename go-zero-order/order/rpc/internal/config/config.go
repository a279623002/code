package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

// DBConfig 数据库配置
type DBConfig struct {
	DataSource      string `json:"DataSource"`
	MaxIdleConns    int    `json:"MaxIdleConns,default=10"`
	MaxOpenConns    int    `json:"MaxOpenConns,default=100"`
	ConnMaxLifetime int    `json:"ConnMaxLifetime,default=3600"`
}

// Config rpc 服务配置
type Config struct {
	zrpc.RpcServerConf
	DB DBConfig
}
