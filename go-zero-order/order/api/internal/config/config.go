package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config api 网关配置
type Config struct {
	rest.RestConf
	OrderRpc zrpc.RpcClientConf
}
