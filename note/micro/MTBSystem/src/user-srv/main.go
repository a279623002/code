package main

import (
	"user-srv/handler"
	pb "user-srv/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user-srv"),
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
	pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserHandler))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
