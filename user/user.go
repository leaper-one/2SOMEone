package main

import (
	"context"
	"errors"
	"fmt"

	// message "github.com/leaper-one/2SOMEone/service/message"
	user "github.com/leaper-one/2SOMEone/service/user"
	"github.com/leaper-one/2SOMEone/util"

	msg_pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/message/v1"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/user/v1"

	"go-micro.dev/v4"
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

type UserService struct{}

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
func (u *UserService) SignUpByPhone(ctx context.Context, req *pb.SignUpByPhoneRequest, response *pb.SignUpByPhoneResponse) error {
	fmt.Printf("SignUpByPhone: %v\n", req)

	// check phone code
	// 创建一个新的服务
	// service := micro.NewService(micro.Name("Greeter.Client"))
	service := micro.NewService(micro.Name("go.micro.service.message.client"))
	// 初始化
	service.Init()

	// 创建 Message 客户端
	// greeter := pb.NewGreeterService("Greeter", service.Client())
	msgr := msg_pb.NewMessageService("go.micro.srv.message", service.Client())

	// 远程调用 Greeter 服务的 Hello 方法
	// rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "学院君"})
	rsp, err := msgr.CheckMessageCode(context.Background(), &msg_pb.CheckMessageCodeRequest{
		Phone: req.Phone,
		Code:  req.Code,
		MsgId: req.MsgId,
	})
	if err != nil {
		return err
	}

	if !rsp.IsMatch {
		return errors.New("code not match")
	}

	err = userService.SignUpByPhone(ctx, req.Phone, req.Password)
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
func (u *UserService) SignInByPhone(ctx context.Context, req *pb.SignInByPhoneRequest, response *pb.SignInByPhoneResponse) error {
	fmt.Printf("SignInByPhone: %v\n", req)
	token, err := userService.Auth(ctx, req.Phone, req.Password)
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
func (u *UserService) GetMe(ctx context.Context, req *pb.GetMeRequest, response *pb.GetMeResponse) error {
	fmt.Printf("GetMe: %v\n", req)
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
func (u UserService) SetInfo(ctx context.Context, req *pb.SetInfoRequest, response *pb.SetInfoResponse) error {
	fmt.Printf("SetInfo: %v\n", req)
	// ctx contains auth_token, param it to get user_id
	user_id, err := util.CheckAuth(ctx)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		return err
	}
	err = userService.SetInfo(ctx, user_id, req.Name, req.Avatar, req.Buid)
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
func (u UserService) GetUserIDByBuid(ctx context.Context, req *pb.GetUserIDByBuidRequest, response *pb.GetUserIDByBuidResponse) error {
	fmt.Printf("GetUserIDByBuid: %v\n", req)
	buser, err := userService.FindByBuid(ctx, req.Buid)
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
