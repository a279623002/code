package main

import (
	"context"
	"fmt"
	"github.com/micro/micro/v3/service"
	proto "service/proto"
)

type Greeter struct {}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Greeting = "hello " + req.Name
	return nil
}

func main() {
	service := service.New(
		service.Name("greeter"),
		)
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}


}
