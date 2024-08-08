package main

import (
	"flag"
	"fmt"
	"net"

	"demo-rpc/internal/config"
	"demo-rpc/internal/server"
	"demo-rpc/internal/svc"
	"demo-rpc/types/demo"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// ctx := svc.NewServiceContext(c)

	// 注册consul服务
	//获取动态接口口
	port, _ := GetFreePort()
	//替换yaml里面的host和端口
	c.ListenOn = fmt.Sprintf("0.0.0.0:%d", port)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		demo.RegisterDemoServer(grpcServer, server.NewDemoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//把服务信息注册到consul
	_ = consul.RegisterService(c.ListenOn, c.Consul)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}

func GetFreePort() (int, error) {
	// 动态获取可用端口
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	fmt.Println(addr.Port) // 0

	l, err := net.Listen("tcp", addr.String())
	if err != nil {
		return 0, err
	}

	return l.Addr().(*net.TCPAddr).Port, nil
}
