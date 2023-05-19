package main

import (
	v1 "client/v1"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"log"
	"time"
)

func main() {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///module"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	gClient := v1.NewModuleClient(conn)

	for {
		time.Sleep(time.Second)
		callGRPC(gClient)
	}
}

func callGRPC(client v1.ModuleClient) {
	ctx, cel := context.WithTimeout(context.Background(), time.Second*5)
	defer cel()
	reply, err := client.GetModule(ctx, &v1.GetModuleRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
