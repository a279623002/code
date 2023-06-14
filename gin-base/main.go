package main

import (
	"edu-admin-api/controller"
	"edu-admin-api/logger"
	"edu-admin-api/pkg/snowflake"
	"edu-admin-api/repository/mysql"
	"edu-admin-api/repository/redis"
	"edu-admin-api/routers"
	"edu-admin-api/settings"
	"fmt"
)

//go web通用配置

func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	//3.初始化MySQL连接-gorm框架
	if err := mysql.InitDB(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql-grom failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	//5.初始化雪花算法
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//6.初始化gin框架内置的校验器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	//7.注册路由-启动服务
	r := routers.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
