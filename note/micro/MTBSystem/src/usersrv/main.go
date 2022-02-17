package main

import (
	"usersrv/handler"
	pb "usersrv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go.micro.srv.user"),
		//service.Name("usersrv"),
		service.Version("latest"),
	)

	// 定义Service动作操作
	srv.Init(
		service.AfterStop(func() error {
			logger.Info("micro.AfterStop test ...")
			return nil
		}),
		service.AfterStart(func() error {
			logger.Info("micro.AfterStart test ...")
			return nil
		}),
	)

	// Register handler
	pb.RegisterUsersrvHandler(srv.Server(), new(handler.UserHandler))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
