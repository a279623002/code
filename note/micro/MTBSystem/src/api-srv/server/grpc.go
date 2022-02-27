package server

import (
	"config"
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
	if paths[1] == "user" {
		if paths[2] == "registAccount" {
			grpc = RegistAccount(c, config.Namespace+config.ServiceNameUser, "User.RegistAccount")
		}
		if paths[2] == "loginAccount" {
			grpc = LoginAccount(c, config.Namespace+config.ServiceNameUser, "User.LoginAccount")
		}
		if paths[2] == "wantScore" {
			grpc = WantScore(c, config.Namespace+config.ServiceNameUser, "User.WantScore")
		}
		if paths[2] == "updateUserProfile" {
			grpc = UpdateUserProfile(c, config.Namespace+config.ServiceNameUser, "User.UpdateUserProfile")
		}

	}
	if grpc == nil {
		err = errors.New("couldn`t found services")
	}
	return
}

func UpdateUserProfile(c *gin.Context, service, endpoint string) *Grpc {
	userName := c.Query("userName")
	userEmail := c.Query("userEmail")
	userPhone := c.Query("userPhone")
	userId,_ := strconv.Atoi(c.Query("userId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &pb.UpdateUserProfileReq{
			UserName:  userName,
			UserEmail: userEmail,
			UserPhone: userPhone,
			UserID:    int64(userId),
		},
		Rsp:      &pb.UpdateUserProfileRsp{},
	}
}
func WantScore(c *gin.Context, service, endpoint string) *Grpc {
	userId,_ := strconv.Atoi(c.Query("userId"))
	movieId,_ := strconv.Atoi(c.Query("movieId"))
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &pb.WantScoreReq{
			UserId:  int64(userId),
			MovieId: int64(movieId),
		},
		Rsp:      &pb.WantScoreRsp{},
	}
}
func LoginAccount(c *gin.Context, service, endpoint string) *Grpc {
	email := c.Query("email")
	password := c.Query("password")
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &pb.LoginAccountReq{
			Email:    email,
			Password: password,
		},
		Rsp:      &pb.LoginAccountRsp{},
	}
}

func RegistAccount(c *gin.Context, service, endpoint string) *Grpc {
	email := c.Query("email")
	userName := c.Query("username")
	password := c.Query("password")
	return &Grpc{
		Service:  service,
		Endpoint: endpoint,
		Req:      &pb.RegistAccountReq{
			Email:    email,
			UserName: userName,
			Password: password,
		},
		Rsp:      &pb.RegistAccountRsp{},
	}
}
