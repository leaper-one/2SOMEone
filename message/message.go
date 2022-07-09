package main

import (
	"context"
	"fmt"

	"github.com/leaper-one/2SOMEone/service/message"
	"github.com/leaper-one/2SOMEone/util"
	pb "github.com/leaper-one/2someone-proto/gen/proto/go/2SOMEone/message/v1"
)

var (
	config     = util.LoadConfig("./config.yaml", &Config{}).(*Config)
	dbc        = util.OpenDB("./message.db")
	msgService = message.NewMsgService(dbc, config.AliMsg.RegionId, config.AliMsg.AccessKeyId, config.AliMsg.AccessKeySecret)
)

type MessageService struct{}

// config struct
// app, alimsg, grpcset
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

// Sent phone message code.
func (m *MessageService) SentMessageCode(ctx context.Context, req *pb.SentMessageCodeRequest, response *pb.SentMessageCodeResponse) error {
	fmt.Printf("SentMessageCode: %v\n", req)
	_, msg_id, err := msgService.SendPhoneCode(ctx, req.Phone)
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		response.MsgId = 0
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	response.MsgId = uint64(msg_id)
	return nil
}

// Check phone message code.
func (m *MessageService) CheckMessageCode(ctx context.Context, req *pb.CheckMessageCodeRequest, response *pb.CheckMessageCodeResponse) error {
	fmt.Printf("CheckMessageCode: %v\n", req)
	is_match, err := msgService.CheckPhoneCode(ctx, req.Phone, req.Code, uint(req.MsgId))
	if err != nil {
		response.Code = util.FAIL
		response.Msg = err.Error()
		response.IsMatch = is_match
		return err
	}
	response.Code = util.SUCCESS
	response.Msg = "success."
	response.IsMatch = is_match
	return nil
}
