package handler

import (
	"config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"net/http"
	"strconv"
	user "user-srv/proto"
)

var (
	serviceUser               = config.Namespace + config.ServiceNameUser
	endpointUpdateUserProfile = "User.UpdateUserProfile"
	endpointWantScore         = "User.WantScore"
	endpointLoginAccount      = "User.LoginAccount"
	endpointRegistAccount     = "User.RegistAccount"
)

func RegistAccount(c *gin.Context) {
	email := c.Query("email")
	userName := c.Query("username")
	password := c.Query("password")

	grpcReq := &user.RegistAccountReq{
		Email:    email,
		UserName: userName,
		Password: password,
	}
	grpcRsp := &user.RegistAccountRsp{}

	req := client.NewRequest(serviceUser, endpointRegistAccount, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

func LoginAccount(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	grpcReq := &user.LoginAccountReq{
		Email:    email,
		Password: password,
	}
	grpcRsp := &user.LoginAccountRsp{}

	req := client.NewRequest(serviceUser, endpointLoginAccount, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

func WantScore(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("userId"))
	movieId, _ := strconv.Atoi(c.Query("movieId"))

	grpcReq := &user.WantScoreReq{
		UserId:  int64(userId),
		MovieId: int64(movieId),
	}
	grpcRsp := &user.WantScoreRsp{}

	req := client.NewRequest(serviceUser, endpointWantScore, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}

func UpdateUserProfile(c *gin.Context) {
	userName := c.Query("userName")
	userEmail := c.Query("userEmail")
	userPhone := c.Query("userPhone")
	userId, _ := strconv.Atoi(c.Query("userId"))

	grpcReq := &user.UpdateUserProfileReq{
		UserName:  userName,
		UserEmail: userEmail,
		UserPhone: userPhone,
		UserID:    int64(userId),
	}
	grpcRsp := &user.UpdateUserProfileRsp{}

	req := client.NewRequest(serviceUser, endpointUpdateUserProfile, grpcReq)

	if err := client.Call(context.Background(), req, grpcRsp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": -1,
			"msg":  err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": grpcRsp,
		})
	}
}
