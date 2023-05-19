package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	options := &redis.ClusterOptions{
		Addrs: []string{"192.168.1.89:7001", "192.168.1.89:7002", "192.168.1.89:7003"},
	}
	client := redis.NewClusterClient(options)
	err := client.Ping()
	fmt.Println(err)
	res := client.Get("shiro").String()
	fmt.Println(res)
}
