package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"cinema-srv/db"
	"cinema-srv/handler"
	pb "cinema-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameCinema),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterCinemaHandler(srv.Server(), new(handler.CinemaServiceExtHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNameCinema +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNameCinema +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
