package example

import (
	"my-redis/configs"
	"time"
	"fmt"
)

func AddAndSmembers() {
	key := "Go-Redis-Set";
	res, _ := configs.RedisDB.SMembers(key).Result()
	if len(res) == 0 {
		fmt.Println("set Go-Redis-Set")
		configs.RedisDB.SAdd(key, 1)
		configs.RedisDB.SAdd(key, 2)
		configs.RedisDB.SAdd(key, 3)
		configs.RedisDB.SAdd(key, 4)
		configs.RedisDB.Expire(key, 60 * time.Second)
		res, _ = configs.RedisDB.SMembers(key).Result()
	}
	fmt.Println(res)
}
