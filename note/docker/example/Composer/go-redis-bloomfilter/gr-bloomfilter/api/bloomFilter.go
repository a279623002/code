package api

import (
	"github.com/gin-gonic/gin"
	"gr-bloomfilter/configs"
	"net/http"
)

func IsExits(c *gin.Context) {
	key := c.Query("key")
	res, err := configs.BFHandler.Exists(c, key)
	if err != nil {
		ExitErr(c, -1, err)
		return
	}
	msg := "exits"
	if !res {
		msg = "no exits"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
	})
}

func Set(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("val")
	res := configs.RedisDB.Set(key, val, -1)
	_, err := res.Result()
	if err != nil {
		ExitErr(c, -1, err)
		return
	}
	err = configs.BFHandler.Add(c, key)
	if err != nil {
		ExitErr(c, -1, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})
}

func Get(c *gin.Context) {
	key := c.Query("key")
	res, err := configs.RedisDB.Get(key).Result()
	if err != nil {
		ExitErr(c, -1, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": res,
	})
}
