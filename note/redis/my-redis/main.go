package main

import (
	"my-redis/configs"
	"my-redis/example"
)

func main() {
	session := configs.IniSessionGet()
	configs.InitDb(session)

	//example.SetAndGet()
	//example.HMSetAndGet()
	//example.PushAndRange()
	//example.AddAndSmembers()
	//example.ZAddAndZRange()
	example.PrintHash()
}