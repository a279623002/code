package main

import (
	"api-srv/server"
)

func main() {
	srv := server.NewApi()
	srv.Run(":8082")
}
