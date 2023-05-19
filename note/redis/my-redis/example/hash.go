package example

import (
	"fmt"
	"my-redis/configs"
	"time"
)

func HMSetAndGet() {
	key := "Go-Redis-Hash"
	res, _ := configs.RedisDB.HGetAll(key).Result()
	if len(res) == 0 {
		fmt.Println("set Go-Redis-Hash")
		data := make(map[string]interface{})
		data["name"] = "shiro"
		data["hobby"] = "hanser"
		configs.RedisDB.HMSet(key, data)
		configs.RedisDB.Expire(key, 60*time.Second)
		res, _ = configs.RedisDB.HGetAll(key).Result()
	}
	fmt.Println(res)

	ret, _ := configs.RedisDB.HGet(key, "hobby").Result()
	fmt.Println(ret)
}
