package main

import (
	"go-mysql/configs"
	"go-mysql/example"
)


func main() {
	mysql := configs.IniMysqlGet()
	configs.InitDb(mysql)

	//example.Select()
	//example.Insert()
	//example.Update()
	//example.Delete()
	example.ToUpdate()
}
