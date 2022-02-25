package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"net/http"
)

type Api struct {
	client *client.Client
	route *gin.Engine
}

func NewApi() *Api {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	r := gin.Default()
	return &Api{
		client: &c,
		route:  r,
	}
}

func (a *Api) Run(port string) {
	a.route.GET("/", DefaultHandler)

	user := a.route.Group("/user")
	{
		user.GET("/selectUser", GrpcHandler)
	}
	a.route.Run(port)
}

var DefaultHandler = func(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}

var GrpcHandler = func(c *gin.Context) {
	grpc, err := GetGrpc(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	req := client.NewRequest(grpc.Service, grpc.Endpoint, grpc.Req)

	if err := client.Call(context.Background(), req, grpc.Rsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpc.Req,
		})
	}
}
