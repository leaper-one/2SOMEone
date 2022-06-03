package main

import (
	"log"
	"net"

	pb "github.com/leaper-one/2someone-proto/gen/golang/account/user"
	"github.com/leaper-one/2SOMEone/util"
	"google.golang.org/grpc"
)

func main() {
	// start server
	config := util.LoadConfig("./config.yaml", &Config{}).(*Config)
	log.Printf("App: %v", config.App.Name)
	lis, err := net.Listen("tcp", config.GrpcSet.EndPoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen in %v", config.GrpcSet.EndPoint)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// pb.RegisterGreeterServer(grpcServer, newServer())
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
