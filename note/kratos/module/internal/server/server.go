package server

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"module/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistry)

func NewRegistry(c *conf.Consul) *consul.Registry {
	cs := api.DefaultConfig()
	cs.Address = c.Address
	cs.Scheme = c.Scheme

	client, err := api.NewClient(cs)
	if err != nil {
		panic(err)
	}
	return consul.New(client, consul.WithHealthCheck(false))
}
