package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"film-srv/db"
	"film-srv/handler"
	pb "film-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameFilm),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterFilmHandler(srv.Server(), new(handler.FilmServiceExtHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNameFilm +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNameFilm +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
