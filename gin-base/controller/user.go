package controller

import (
	"edu-admin-api/models"
	"edu-admin-api/pkg/jwt"
	"edu-admin-api/repository/mysql"
	"edu-admin-api/service"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 用户注册
func SignUpHandler(c *gin.Context) {
	//1.获取参数，参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.业务逻辑处理
	if err := service.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)
}

// 用户登录
func LoginHandler(c *gin.Context) {
	//1.获取参数，参数校验
	p := make(map[string]interface{})
	p["username"] = c.PostForm("username")
	p["password"] = c.PostForm("password")
	if c.PostForm("username") == "" || c.PostForm("password") == "" {
		ResponseError(c, CodeInvalidParams)
		return
	}
	//2.业务逻辑处理
	u, err := service.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应-accessToken
	aToken, rToken, _ := jwt.GenToken(u.UserID)
	ResponseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       u.UserID,
		"username":     u.Username,
	})

}
