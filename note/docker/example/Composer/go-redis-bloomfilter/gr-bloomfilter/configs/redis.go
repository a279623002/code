package configs

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strings"
)

var RedisDB *redis.Client

var BFHandler *BloomFilter

func InitDb(cfg *SessionConfig) {
	// 设置环境变量

	os.Setenv("REDIS_HOST", cfg.Host)

	os.Setenv("REDIS_PORT", cfg.Port)
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		DB:       cfg.Select,
		Password: cfg.Authstring,
	})

	BFHandler = &BloomFilter{
		Bucket:    cfg.BFBucket,
		HashFuncs: strings.Split(cfg.BFFunc, ","),
	}
}
