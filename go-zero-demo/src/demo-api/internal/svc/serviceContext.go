package svc

import (
	"demo-api/internal/config"
	"demo-rpc/democlient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	DemoClient democlient.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		DemoClient: democlient.NewDemo(zrpc.MustNewClient(c.DemoRpc)),
	}
}
