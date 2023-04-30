package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExitErr(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  err.Error(),
	})
	//终止后续接口调用，不执行接口里后续代码
	c.Abort()
}
