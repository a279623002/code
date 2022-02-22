package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	pb "user-srv/proto"
	"user-srv/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go.micro.user-srv1"),
		service.Version("latest"),
	)

	// Register handler
	err := pb.RegisterUserSrvHandler(srv.Server(), new(handler.UserHandler))
	if err != nil {
		panic(err)
	}

	srv.Init(
		service.AfterStop(func() error {
			logger.Info("user-srv stot")
			return nil
		}),
		service.AfterStart(func() error {
			logger.Info("user-srv start")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
