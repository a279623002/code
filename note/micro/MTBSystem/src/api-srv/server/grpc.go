package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	pb "user-srv/proto"
)

type Grpc struct {
	Service  string
	Endpoint string
	Req      interface{}
	Rsp      interface{}
}

func GetGrpc(c *gin.Context) (grpc *Grpc, err error) {
	path := c.Request.URL.Path
	paths := strings.Split(path, "/")
	if len(paths) != 3 {
		return &Grpc{}, errors.New("bad request")
	}
	if paths[1] == "user" && paths[2] == "selectUser" {
		grpc = selectUser(c, "go.micro.user", "User.SelectUser")
	}
	if grpc == nil {
		err = errors.New("couldn`t found services")
	}
	return
}

func selectUser(c *gin.Context, service, endpoint string) *Grpc {
	paramsId := c.Query("id")
	id, _ := strconv.Atoi(paramsId)
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &pb.SelectUserReq{Id: int64(id)},
		Rsp:      &pb.SelectUserResp{},
	}
}
