package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	proto "service/proto"
)

type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		)
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}


}
