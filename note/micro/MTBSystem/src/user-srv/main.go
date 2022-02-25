package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"user-srv/handler"
	pb "user-srv/proto"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.user"),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			pb.RegisterUserHandler(srv.Server(), new(handler.UserHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println("user-srv start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println("user-srv stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}