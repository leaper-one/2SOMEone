package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/leaper-one/2SOMEone/grpc/user"
	"github.com/leaper-one/2SOMEone/util"
	"google.golang.org/grpc"
)

const (
	HTTPLISTEN = ":8081" //http端口
	GRPCLISTEN = ":9090" //grpc端口
)

func main() {
	// start server
	config := util.LoadConfig("./config.yaml", &Config{}).(*Config)
	log.Printf("App: %v", config.App.Name)
	//lis, err := net.Listen("tcp", config.GrpcSet.EndPoint)
	lis, err := net.Listen("tcp", GRPCLISTEN)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen in %v", config.GrpcSet.EndPoint)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// pb.RegisterGreeterServer(grpcServer, newServer())
	pb.RegisterUserServiceServer(grpcServer, &UserService{})

	/*2022-7.4 添加http接口监听grpc服务
	 */
	go grpcServer.Serve(lis) //启动grpc服务

	mux := runtime.NewServeMux()
	pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, GRPCLISTEN, []grpc.DialOption{grpc.WithInsecure()})
	http.ListenAndServe(HTTPLISTEN, mux) //http 8081端口
	/**/
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}
}
