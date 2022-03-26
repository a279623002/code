package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"place-srv/db"
	"place-srv/handler"
	pb "place-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNamePlace),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterPlaceHandler(srv.Server(), new(handler.PlaceServiceExtHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNamePlace +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNamePlace +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
