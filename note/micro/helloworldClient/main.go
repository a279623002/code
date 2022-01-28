package main

import (
	"context"
	"fmt"
	"github.com/micro/micro/v3/service"
	proto "helloworld/proto"
)

func main() {
	server := service.New()
	client := proto.NewHelloworldService("helloworld", server.Client())

	rsp, err := client.Call(context.Background(), &proto.Request{Name:"client"})
	if err != nil {
		fmt.Println("Error Calling helloworld: ", err)
		return
	}

	fmt.Println("Response: ", rsp.Msg)
}
