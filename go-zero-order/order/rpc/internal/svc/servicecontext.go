package svc

import (
	"go-zero-service/common/gorm"
	"go-zero-service/order/rpc/internal/config"
	"go-zero-service/order/rpc/internal/model"
)

// ServiceContext rpc 服务上下文
type ServiceContext struct {
	Config     config.Config
	OrderModel *model.OrderModel
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

	orderModel := model.NewOrderModel(db)
	if err := orderModel.AutoMigrate(); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		OrderModel: orderModel,
	}
}
