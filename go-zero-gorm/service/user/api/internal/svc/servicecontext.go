package svc

import (
	"go-zero-gorm/service/user/api/internal/config"
	model "go-zero-gorm/service/user/model"
	user "go-zero-gorm/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	UserRpc   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
