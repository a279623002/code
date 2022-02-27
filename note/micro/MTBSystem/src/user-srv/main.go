package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"user-srv/db"
	"user-srv/handler"
	pb "user-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameUser),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterUserHandler(srv.Server(), new(handler.UserHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNameUser +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNameUser +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}