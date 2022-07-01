package main

import (
	"context"
	"fmt"

	// message "github.com/leaper-one/2SOMEone/service/message"
	user "github.com/leaper-one/2SOMEone/service/user"
	"github.com/leaper-one/2SOMEone/util"

	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/user/v1"
)

// const (
// SUCCESS = 200
// FAIL    = 500
// )

var (
	// config      = util.LoadConfig("./config.yaml", &Config{}).(*Config)
	dbc         = util.OpenDB("./user.db")
	userService = user.NewUserService(dbc)
	// msgService  = message.NewMsgService(dbc, config.AliMsg.RegionId, config.AliMsg.AccessKeyId, config.AliMsg.AccessKeySecret)
)

type UserService struct {
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
	Registry struct {
		Address string `yaml:"address"`
	}
}

// Sign up by phone
func (u *UserService) SignUpByPhone(ctx context.Context, request *pb.SignUpByPhoneRequest, response *pb.SignUpByPhoneResponse) error {
	fmt.Printf("SignUpByPhone: %v\n", request)
	err := userService.SignUpByPhone(ctx, request.Phone, request.Code, request.Password, uint(request.MsgId))
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	return nil
}

// Sign in by phone
func (u *UserService) SignInByPhone(ctx context.Context, request *pb.SignInByPhoneRequest, response *pb.SignInByPhoneResponse) error {
	fmt.Printf("SignInByPhone: %v\n", request)
	token, err := userService.Auth(ctx, request.Phone, request.Password)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		response.Token = ""
		return err
	}

	response.Code = util.SUCCESS
	response.Msg = "success."
	response.Token = token
	return nil
}

// Get me
func (u *UserService) GetMe(ctx context.Context, request *pb.GetMeRequest, response *pb.GetMeResponse) error {
	fmt.Printf("GetMe: %v\n", request)
	// ctx contains auth_token, param it to get user_id
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	user, err := userService.GetMe(ctx, user_id)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	response.User = &pb.BasicUser{
		UserId: user.UserID,
		UserInfo: &pb.UserInfo{
			Name:   user.Name,
			Avatar: user.Avatar,
			Phone:  user.Phone,
			Email:  user.Email,
		},
	}
	return nil
}
// Set info
func (u UserService) SetInfo(ctx context.Context, request *pb.SetInfoRequest, response *pb.SetInfoResponse) error {
	fmt.Printf("SetInfo: %v\n", request)
	// ctx contains auth_token, param it to get user_id
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	err = userService.SetInfo(ctx, user_id, request.Name, request.Avatar, request.Buid)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	return nil
}

// Get user_id by buid
func (u UserService) GetUserIDByBuid(ctx context.Context, request *pb.GetUserIDByBuidRequest, response *pb.GetUserIDByBuidResponse) error {
	fmt.Printf("GetUserIDByBuid: %v\n", request)
	buser, err := userService.FindByBuid(ctx, request.Buid)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	response.UserId = buser.UserID
	return nil
}
