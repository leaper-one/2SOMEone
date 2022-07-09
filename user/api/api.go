package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/user/v1"

	"go-micro.dev/v4"
	api "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/errors"
)

type User struct {
	Client pb.UserService
}

func (u *User) SignUpByPhone(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.SignUpByPhone API 请求")

	// phone
	phone, ok := req.Get["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Phone is required")
	}
	// code
	code, ok := req.Get["code"]
	if !ok || len(code.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Code is required")
	}
	// password
	password, ok := req.Get["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Password is required")
	}
	// msg_id
	msg_id, ok := req.Get["msg_id"]
	if !ok || len(msg_id.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Msg_id is required")
	}

	i, _ := strconv.ParseUint(strings.Join(msg_id.Values, " "), 10, 64)
	// 将参数交由底层服务处理
	response, err := u.Client.SignUpByPhone(ctx, &pb.SignUpByPhoneRequest{
		Phone:    strings.Join(phone.Values, " "),
		Code:     strings.Join(code.Values, " "),
		Password: strings.Join(password.Values, " "),
		MsgId: i,
	})

	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func (u *User) SignInByPhone(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.SignInByPhone API 请求")

	// phone
	phone, ok := req.Get["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Phone is required")
	}

	// password
	password, ok := req.Get["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Password is required")
	}

	// 将参数交由底层服务处理
	response, err := u.Client.SignInByPhone(ctx, &pb.SignInByPhoneRequest{
		Phone:    strings.Join(phone.Values, " "),
		Password: strings.Join(password.Values, " "),
	})

	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func (u *User) GetMe(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.GetMe API 请求")

	// 将参数交由底层服务处理
	response, err := u.Client.GetMe(ctx, &pb.GetMeRequest{})

	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func (u *User) SetInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.SetInfo API 请求")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Name is required")
	}
	avatar, ok := req.Get["avatar"]
	if !ok || len(avatar.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Avatar is required")
	}

	buid, ok := req.Get["buid"]
	if !ok || len(buid.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Buid is required")
	}

	// 将参数交由底层服务处理
	response, err := u.Client.SetInfo(ctx, &pb.SetInfoRequest{
		Name:    strings.Join(name.Values, " "),
		Avatar:  strings.Join(avatar.Values, " "),
		Buid: strings.Join(buid.Values, " "),
	})

	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func(u *User) GetUserIDByBuid(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 User.GetUserIDByBuid API 请求")

	buid, ok := req.Get["buid"]
	if !ok || len(buid.Values) == 0 {
		return errors.BadRequest("go.micro.api.user", "Buid is required")
	}
	i, _ := strconv.ParseInt(strings.Join(buid.Values, " "), 10, 64)

	// 将参数交由底层服务处理
	response, err := u.Client.GetUserIDByBuid(ctx, &pb.GetUserIDByBuidRequest{
		Buid: i,
	})

	if err != nil {
		return err
	}

	// 处理成功，则返回处理结果
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	// 创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.api.message"),
	)

	// 解析命令行参数
	service.Init()

	// 将请求转发给底层 go.micro.srv.message 服务处理
	service.Server().Handle(
		service.Server().NewHandler(
			// &Say{Client: hello.NewGreeterService("go.micro.srv.greeter", service.Client())},
			// &Msg{Client: pb.NewMessageService("go.micro.srv.message", service.Client())},
			&User{Client: pb.NewUserService("go.micro.srv.user", service.Client())},
		),
	)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
