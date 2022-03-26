package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"log"
	"comment-srv/db"
	"comment-srv/handler"
	pb "comment-srv/proto"
	"config"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(config.Namespace+config.ServiceNameComment),
		micro.Version("latest"),
	)

	srv.Init(
		micro.Action(func(c *cli.Context) error {
			db.Init(config.MysqlDSN)
			pb.RegisterCommentHandler(srv.Server(), new(handler.CommentServiceExtHandler))
			return nil
		}),
		micro.AfterStart(func() error {

			log.Println(config.ServiceNameComment +"start")
			return nil
		}),
		micro.AfterStop(func() error {

			log.Println(config.ServiceNameComment +"stop")
			return nil
		}),
	)


	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
