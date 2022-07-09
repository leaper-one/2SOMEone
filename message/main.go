package main

import (
	"log"

	"github.com/leaper-one/2SOMEone/util"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/message/v1"

	// "google.golang.org/grpc"

	etcd "github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	// start server
	config := util.LoadConfig("./config.yaml", &Config{}).(*Config)
	log.Printf("App: %v", config.App.Name)

	// 配置etcd
	reg := etcd.NewRegistry(registry.Addrs(config.Registry.Address))

	service := micro.NewService(
		micro.Name("go.micro.srv.message"),
		micro.Address(config.EndPoint.GrpcEndpoint),
		micro.Version("v1"),
		micro.Registry(reg),
	)

	// 初始化，解析命令行参数
	service.Init()
	pb.RegisterMessageHandler(service.Server(), new(MessageService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
