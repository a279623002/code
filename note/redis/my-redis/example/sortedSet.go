package example

import (
	"github.com/go-redis/redis"
	"my-redis/configs"
	"time"
	"fmt"
)

func ZAddAndZRange() {
	key := "Go-Redis-SortedSet";
	res, _ := configs.RedisDB.ZRange(key, 0, -1).Result()
	if len(res) == 0 {
		fmt.Println("set Go-Redis-SortedSet")
		configs.RedisDB.ZAdd(key, redis.Z{float64(91), "score1"})
		configs.RedisDB.ZAdd(key, redis.Z{float64(92), "score2"})
		configs.RedisDB.ZAdd(key, redis.Z{float64(93), "score3"})
		configs.RedisDB.ZAdd(key, redis.Z{float64(94), "score4"})
		configs.RedisDB.Expire(key, 60 * time.Second)
		res, _ = configs.RedisDB.ZRange(key, 0, -1).Result()
	}
	fmt.Println(res)

	ret, _ := configs.RedisDB.ZRangeWithScores(key, 0, -1).Result()
	fmt.Println(ret)
	for _, v := range ret {
		if v.Score > float64(92) {
			fmt.Println(v.Member)
		}
	}
}
