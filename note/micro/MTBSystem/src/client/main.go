package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	pb "user-srv/proto"
)

func main() {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	// Create new request to service go.micro.srv.example, method Example.Call
	req := c.NewRequest("go.micro.user", "User.SelectUser", &pb.SelectUserReq{
		Id: 2233,
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		//"X-User-Id": "john",
		//"X-From-Id": "script",
	})

	rsp := &pb.SelectUserResp{}

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call: rsp:", rsp.User)
}