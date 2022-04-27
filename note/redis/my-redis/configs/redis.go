package configs

import "github.com/go-redis/redis"

var RedisDB *redis.Client

func InitDb(configs *SessionConfig) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     configs.Host,
		DB:       configs.Select,
		Password: configs.Authstring,
	})
}
