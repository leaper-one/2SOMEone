package main

import (
	"2SOMEone/service"
	"2SOMEone/util"
	"context"
	"fmt"
	"log"
	"net"

	pb "2SOMEone/grpc/user"

	"google.golang.org/grpc"
)

const (
	SUCCESS = 200
	FAIL    = 500
)

var (
	dbc         = util.OpenDB("./2-some-one.db")
	userService = service.NewUserService(dbc)
	noteService = service.NewNoteService(dbc)
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

// Sent phone message code.
func (s *UserService) SentMessageCode(ctx context.Context, in *pb.SentMessageCodeRequest) (*pb.SentMessageCodeResponse, error) {
	fmt.Printf("SentMessageCode: %v\n", in)
	code, err := userService.SendPhoneCode(ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	fmt.Printf("code: %v\n", code)
	return &pb.SentMessageCodeResponse{Code: SUCCESS, Msg: "success."}, nil
}

func (s *UserService) SignUpByPhone(ctx context.Context, in *pb.SignUpByPhoneRequest) (*pb.SignUpByPhoneResponse, error) {
	fmt.Printf("SignUpByPhone: %v\n", in)
	err := userService.SignUpByPhone(ctx, in.Phone, in.Code, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpByPhoneResponse{Code: SUCCESS, Msg: "success."}, nil
}

func (s *UserService) SignInByPhone(ctx context.Context, in *pb.SignInByPhoneRequest) (*pb.SignInByPhoneResponse, error) {
	fmt.Printf("SignInByPhone: %v\n", in)
	token, err := userService.Auth(ctx, in.Phone, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.SignInByPhoneResponse{Code: SUCCESS, Msg: "success.", Token: token}, nil
}

func main() {

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
