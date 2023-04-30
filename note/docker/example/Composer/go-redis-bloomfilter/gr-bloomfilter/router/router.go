package router

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"gr-bloomfilter/api"
	"gr-bloomfilter/configs"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "hello",
	})
}

func BfBucketHandler(c *gin.Context) {
	key := c.Query("key")
	// 防止key与bucket相同
	cfg, err := configs.IniSessionGet()
	if err != nil {
		api.ExitErr(c, -1, err)
		return
	}
	if key == cfg.BFBucket {
		api.ExitErr(c, -1, errors.New("error"))
		return
	}

	c.Next()
}

func Run(ctx context.Context) {
	r := gin.Default()
	r.Use(BfBucketHandler)
	r.GET("/", DefaultHandler)
	user := r.Group("/bloomFilter")
	{
		user.GET("/isExists", api.IsExits)
		user.POST("/set", api.Set)
		user.GET("/get", api.Get)
	}
	r.Run(configs.Server.Port)
}
