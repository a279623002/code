package routers

import (
	"edu-admin-api/controller"
	"edu-admin-api/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v2 := r.Group("/api/v1")
	v2.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "GET",
		})
	})
	v2.POST("/signup", controller.SignUpHandler)
	v2.POST("/login", controller.LoginHandler)
	v2.Use(controller.JWTAuthMiddleware())
	{
		v2.POST("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "POST",
			})
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
