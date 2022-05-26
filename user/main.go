package main

import (
	"log"
	"net"

	pb "2SOMEone/grpc/user"
	"google.golang.org/grpc"
)

func main() {
	// start server
	config := loadConfig("./config.yaml")
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
