package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	proto "service/proto"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name:"shiro"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.String())
}