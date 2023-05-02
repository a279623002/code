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

// @Summary 用户注册
// @Tags 用户中心
// @Description 用户注册
// @Accept json
// @Produce json
// @Param email query string true "邮箱"
// @Param username query string true "账号"
// @Param password query string true "密码"
// @Success 1 {object} model.Response "success"
// @Router /user/registAccount [post]
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

// @Summary 用户登录
// @Tags 用户中心
// @Description 用户登录
// @Accept json
// @Produce json
// @Param email query string true "邮箱"
// @Param password query string true "密码"
// @Success 1 {object} model.Response "success"
// @Router /user/loginAccount [post]
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

// @Summary 电影评分
// @Tags 用户中心
// @Description 评分（需要下单后才能评分）
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Param movieId query int true "电影id"
// @Success 1 {object} model.Response "success"
// @Router /user/wantScore [post]
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

// @Summary 用户更新
// @Tags 用户中心
// @Description 用户更新
// @Accept json
// @Produce json
// @Param userId query int true "用户id"
// @Param email query string true "邮箱"
// @Param username query string true "账号"
// @Param phone query string true "手机"
// @Success 1 {object} model.Response "success"
// @Router /user/updateUserProfile [post]
func UpdateUserProfile(c *gin.Context) {
	userName := c.Query("username")
	userEmail := c.Query("email")
	userPhone := c.Query("phone")
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
