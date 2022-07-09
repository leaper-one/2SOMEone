package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/leaper-one/2SOMEone/util"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/message/v1"

	"go-micro.dev/v4"
	api "go-micro.dev/v4/api/proto"
	"go-micro.dev/v4/errors"
)

type Msg struct {
	Client pb.MessageService
}

func (m *Msg) SentMessageCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Msg.SentMessageCode API 请求")

	phone, ok := req.Get["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest("go.micro.api.message", "Phone is required")
	}

	// 将参数交由底层服务处理
	response, err := m.Client.SentMessageCode(ctx, &pb.SentMessageCodeRequest{
		Phone: strings.Join(phone.Values, " "),
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

func (m *Msg) CheckMessageCode(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到 Msg.CheckMessageCode API 请求")

	phone, ok := req.Get["phone"]
	if !ok || len(phone.Values) == 0 {
		return errors.BadRequest("go.micro.api.message", "Phone is required")
	}

	code, ok := req.Get["code"]
	if !ok || len(code.Values) == 0 {
		return errors.BadRequest("go.micro.api.message", "Code is required")
	}

	msg_id, ok := req.Get["msg_id"]
	if !ok || len(msg_id.Values) == 0 {
		return errors.BadRequest("go.micro.api.message", "MsgId is required")
	}
	i, _ := strconv.ParseUint(strings.Join(msg_id.Values, " "), 10, 64)
	// 将参数交由底层服务处理
	response, err := m.Client.CheckMessageCode(ctx, &pb.CheckMessageCodeRequest{
		Phone: strings.Join(phone.Values, " "),
		Code:  strings.Join(code.Values, " "),
		// MsgId: uint64(msg_id.Values[0].(float64)),
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

type Config struct {
	App struct {
		Name string `yaml:"name"`
	}
	EndPoint struct {
		ApiEndpoint string `yaml:"api_endpoint"`
	}
}

func main() {
	config := util.LoadConfig("./config.yaml", &Config{}).(*Config)
	// 创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.api.message"),
		micro.Address(config.EndPoint.ApiEndpoint),
	)

	// 解析命令行参数
	service.Init()

	// 将请求转发给底层 go.micro.srv.message 服务处理
	service.Server().Handle(
		service.Server().NewHandler(
			// &Say{Client: hello.NewGreeterService("go.micro.srv.greeter", service.Client())},
			&Msg{Client: pb.NewMessageService("go.micro.srv.message", service.Client())},
		),
	)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
