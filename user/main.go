package main

import (
	"log"
	// "net"

	"github.com/leaper-one/2SOMEone/util"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/user/v1"

	// "google.golang.org/grpc"

	"go-micro.dev/v4"
	// "go-micro.dev/v4/registry"
	// etcd "go-micro.dev/v4/registry"
	// "github.com/micro/go-micro"
	// "github.com/micro/go-micro/registry"
	// "github.com/micro/go-micro/registry/etcd"
)

func main() {
	// start server
	config := util.LoadConfig("./config.yaml", &Config{}).(*Config)
	log.Printf("App: %v", config.App.Name)

	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("v1"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		// micro.Registry(etcd.NewRegistry(
			// etcd address
			// registry.Addrs(config.Registry.Address),
		// )),
		// micro.Registry(registry.),
	)

	service.Init()

	pb.RegisterUserServiceHandler(service.Server(), new(UserService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	// lis, err := net.Listen("tcp", config.GrpcSet.EndPoint)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Printf("listen in %v", config.GrpcSet.EndPoint)
	// var opts []grpc.ServerOption
	// grpcServer := grpc.NewServer(opts...)
	// // pb.RegisterGreeterServer(grpcServer, newServer())
	// pb.RegisterUserServiceServer(grpcServer, &UserService{})
	// err = grpcServer.Serve(lis)
	// if err != nil {
	// 	return
	// }
}
