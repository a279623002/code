package svc

import (
	"go-zero-gorm/service/user/model"
	"go-zero-gorm/service/user/rpc/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	logger "go-zero-gorm/logger"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: NewUserModel(conn),
	}
}

func NewUserModel(sqlConn sqlx.SqlConn) model.UserModel {
	db, err := sqlConn.RawDB()
	if err != nil {
		logx.Error("conn database fail")
		return nil
	}
	newLogger := logger.NewZeroLog(glog.Config{
		SlowThreshold:             time.Second,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  glog.Info,
	})
	g, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{Logger: newLogger})
	if err != nil {
		logx.Error("open gorm fail")
		return nil
	}
	return model.NewUserModel(sqlConn, g)
}
