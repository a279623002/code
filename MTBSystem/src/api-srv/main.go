package main

import (
	"api-srv/server"
)

// @title 影院订票系统
// @version 1.0
// @description 影院订票系统.
// @host 127.0.0.1:8082
// @BasePath /
func main() {
	srv := server.NewApi()
	srv.Run(":8082")
}
