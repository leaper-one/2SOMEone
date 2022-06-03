package main

import (
	"context"
	"fmt"

	"github.com/leaper-one/2SOMEone/service"
	"github.com/leaper-one/2SOMEone/util"

	pb "github.com/leaper-one/2someone-proto/gen/golang/account/user"
)

const (
	SUCCESS = 200
	FAIL    = 500
)

var (
	config      = util.LoadConfig("./config.yaml", &Config{}).(*Config)
	dbc         = util.OpenDB("./user.db")
	userService = service.NewUserService(dbc)
	msgService  = service.NewMsgService(dbc, config.AliMsg.RegionId, config.AliMsg.AccessKeyId, config.AliMsg.AccessKeySecret)
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}
type Config struct {
	App struct {
		Name string `yaml:"name"`
	}
	AliMsg struct {
		RegionId        string `yaml:"region_id"`
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
	}
	GrpcSet struct {
		EndPoint string `yaml:"end_point"`
	}
}

// Sent phone message code.
func (s *UserService) SentMessageCode(ctx context.Context, in *pb.SentMessageCodeRequest) (*pb.SentMessageCodeResponse, error) {
	fmt.Printf("SentMessageCode: %v\n", in)
	_, msg_id, err := msgService.SendPhoneCode(ctx, in.Phone)
	if err != nil {
		return &pb.SentMessageCodeResponse{Code: FAIL, Msg: err.Error()}, err
	}
	return &pb.SentMessageCodeResponse{Code: SUCCESS, Msg: "success.", MsgId: msg_id}, nil
}

func (s *UserService) SignUpByPhone(ctx context.Context, in *pb.SignUpByPhoneRequest) (*pb.SignUpByPhoneResponse, error) {
	fmt.Printf("SignUpByPhone: %v\n", in)
	err := userService.SignUpByPhone(ctx, in.Phone, in.Code, in.Password, uint(in.MsgId))
	if err != nil {
		return &pb.SignUpByPhoneResponse{Code: FAIL, Msg: err.Error()}, err
	}
	return &pb.SignUpByPhoneResponse{Code: SUCCESS, Msg: "success."}, nil
}

func (s *UserService) SignInByPhone(ctx context.Context, in *pb.SignInByPhoneRequest) (*pb.SignInByPhoneResponse, error) {
	fmt.Printf("SignInByPhone: %v\n", in)
	token, err := userService.Auth(ctx, in.Phone, in.Password)
	if err != nil {
		return &pb.SignInByPhoneResponse{Code: FAIL, Msg: err.Error(), Token: ""}, err
	}
	return &pb.SignInByPhoneResponse{Code: SUCCESS, Msg: "success.", Token: token}, nil
}

// Get current user infomation
func (s *UserService) GetMe(ctx context.Context, in *pb.GetMeRequest) (*pb.GetMeResponse, error) {
	fmt.Printf("GetMe: %v\n", in)
	// ctx contains auth_token, param it to get user_id
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		return &pb.GetMeResponse{Code: FAIL, Msg: err.Error()}, err
	}
	user, err := userService.GetMe(ctx, user_id)
	if err != nil {
		return &pb.GetMeResponse{Code: FAIL, Msg: err.Error()}, err
	}

	return &pb.GetMeResponse{Code: SUCCESS, Msg: "success.", User: &pb.BasicUser{
		UserId: user.UserID,
		UserInfo: &pb.UserInfo{
			Name:   user.Name,
			Phone:  user.Phone,
			Avatar: user.Avatar,
			Email:  user.Email,
		},
	}}, nil
}

func (s *UserService) SetInfo(ctx context.Context, in *pb.SetInfoRequest) (*pb.SetInfoResponse, error) {
	fmt.Printf("SetInfo: %v\n", in)
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		return &pb.SetInfoResponse{Code: FAIL, Msg: err.Error()}, err
	}
	err = userService.SetInfo(ctx, user_id, in.Name, in.Avatar, in.Buid)
	if err != nil {
		return &pb.SetInfoResponse{Code: FAIL, Msg: err.Error()}, err
	}
	return &pb.SetInfoResponse{Code: SUCCESS, Msg: "success."}, nil
}

func (s *UserService) GetUserIDByBuid(ctx context.Context, in *pb.GetUserIDByBuidRequest) (*pb.GetUserIDByBuidResponse, error) {
	fmt.Printf("in: %v\n", in)
	buser, err := userService.FindByBuid(ctx, in.Buid)
	if err != nil {
		return &pb.GetUserIDByBuidResponse{Code: FAIL, Msg: err.Error()}, err
	}
	return &pb.GetUserIDByBuidResponse{Code: SUCCESS, Msg: "success.", UserId: buser.UserID}, nil

}