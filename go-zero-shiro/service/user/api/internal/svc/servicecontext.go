package svc

import (
	"go-zero-shiro/service/user/api/internal/config"
	model "go-zero-shiro/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	user "go-zero-shiro/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.GzsUserModel
	UserRpc   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewGzsUserModel(conn, nil),
		UserRpc:   user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
