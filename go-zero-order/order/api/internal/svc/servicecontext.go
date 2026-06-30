package svc

import (
	"go-zero-service/order/api/internal/config"
	"go-zero-service/order/rpc/orderclient"
)

// ServiceContext api 网关服务上下文
type ServiceContext struct {
	Config    config.Config
	OrderRpc  orderclient.Order
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	client, err := orderclient.NewClient(c.OrderRpc)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:   c,
		OrderRpc: client,
	}
}
