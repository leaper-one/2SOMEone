// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

package messageservice

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/message/rpc/types/message"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckMessageCodeRequest  = message.CheckMessageCodeRequest
	CheckMessageCodeResponse = message.CheckMessageCodeResponse
	SentMessageCodeRequest   = message.SentMessageCodeRequest
	SentMessageCodeResponse  = message.SentMessageCodeResponse

	MessageService interface {
		//  向手机号发送验证码
		SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...grpc.CallOption) (*SentMessageCodeResponse, error)
		//  校验验证码
		CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...grpc.CallOption) (*CheckMessageCodeResponse, error)
	}

	defaultMessageService struct {
		cli zrpc.Client
	}
)

func NewMessageService(cli zrpc.Client) MessageService {
	return &defaultMessageService{
		cli: cli,
	}
}

//  向手机号发送验证码
func (m *defaultMessageService) SentMessageCode(ctx context.Context, in *SentMessageCodeRequest, opts ...grpc.CallOption) (*SentMessageCodeResponse, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.SentMessageCode(ctx, in, opts...)
}

//  校验验证码
func (m *defaultMessageService) CheckMessageCode(ctx context.Context, in *CheckMessageCodeRequest, opts ...grpc.CallOption) (*CheckMessageCodeResponse, error) {
	client := message.NewMessageServiceClient(m.cli.Conn())
	return client.CheckMessageCode(ctx, in, opts...)
}
