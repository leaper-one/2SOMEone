package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/leaper-one/2SOMEone/grpc/user"
	"google.golang.org/grpc"
)

func main() {
	// start server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// pb.RegisterGreeterServer(grpcServer, newServer())
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
