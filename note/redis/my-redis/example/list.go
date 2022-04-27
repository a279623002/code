package example

import (
	"my-redis/configs"
	"time"
	"fmt"
)

func PushAndRange() {
	key := "Go-Redis-List";
	res, _ := configs.RedisDB.LRange(key, 0, -1).Result()
	if len(res) == 0 {
		fmt.Println("set Go-Redis-List")
		configs.RedisDB.LPush(key, 1)
		configs.RedisDB.LPush(key, 2)
		configs.RedisDB.LPush(key, 3)
		configs.RedisDB.LPush(key, 4)
		configs.RedisDB.Expire(key, 60 * time.Second)
		res, _ = configs.RedisDB.LRange(key, 0, -1).Result()
	}
	fmt.Println(res)
}
