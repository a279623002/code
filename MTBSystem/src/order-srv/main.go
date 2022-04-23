package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"order-srv/db"
	"order-srv/handler"
	pb "order-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameOrder),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterOrderHandler(srv.Server(), new(handler.OrderServiceExtHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNameOrder +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNameOrder +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
