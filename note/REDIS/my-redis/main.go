package main

import (
	"fmt"
	"my-redis/configs"
)

func main() {
	session := configs.IniSessionGet()
	configs.InitDb(session)

	res, _ := configs.RedisDB.SCard("wego_goods_id").Result()
	fmt.Println(res)
}