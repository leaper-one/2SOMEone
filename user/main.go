package main

import (
	"log"

	"github.com/leaper-one/2SOMEone/util"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/user/v1"

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
		micro.Name("go.micro.srv.user"),
		micro.Version("v1"),
		micro.Address(config.EndPoint.GrpcEndpoint),
		micro.Registry(reg),
	)

	// 初始化，解析命令行参数
	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), new(UserService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
