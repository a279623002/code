package main

import (
	"context"
	"gr-bloomfilter/configs"
	"gr-bloomfilter/router"
)

func main() {
	session, err := configs.IniSessionGet()
	if err != nil {
		panic(err)
	}
	configs.InitDb(session)
	server, err := configs.IniServerGet()
	if err != nil {
		panic(err)
	}
	configs.InitServer(server)

	router.Run(context.Background())
}
