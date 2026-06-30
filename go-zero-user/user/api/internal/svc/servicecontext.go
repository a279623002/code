package svc

import (
	"go-zero-order/order/rpc/orderclient"
	"go-zero-user/common/gorm"
	"go-zero-user/user/api/internal/config"
	"go-zero-user/user/api/internal/model"
)

// ServiceContext user api 服务上下文
type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
	OrderRpc  orderclient.Order
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.New(gorm.Config{
		DataSource:      c.DB.DataSource,
		MaxIdleConns:    c.DB.MaxIdleConns,
		MaxOpenConns:    c.DB.MaxOpenConns,
		ConnMaxLifetime: c.DB.ConnMaxLifetime,
	})
	if err != nil {
		panic(err)
	}

	userModel := model.NewUserModel(db)
	if err := userModel.AutoMigrate(); err != nil {
		panic(err)
	}

	orderClient, err := orderclient.NewClient(c.OrderRpc)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		UserModel: userModel,
		OrderRpc:  orderClient,
	}
}
