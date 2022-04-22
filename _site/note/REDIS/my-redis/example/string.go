package example

import (
	"fmt"
	"my-redis/configs"
	"time"
)

func SetAndGet() {
	key := "Go-Redis-String";
	res, _ := configs.RedisDB.Get(key).Result()
	if res == "" {
		fmt.Println("set Go-Redis-String")
		res = "2233"
		configs.RedisDB.Set(key, "2233", 60*time.Second)
	}
	fmt.Println(res)
}
