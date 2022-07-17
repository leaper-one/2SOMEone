// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

package server

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/message/rpc/message/internal/logic"
	"github.com/leaper-one/2SOMEone/2someone/message/rpc/message/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/message/rpc/message/types/message"
)

type MessageServer struct {
	svcCtx *svc.ServiceContext
	message.UnimplementedMessageServer
}

func NewMessageServer(svcCtx *svc.ServiceContext) *MessageServer {
	return &MessageServer{
		svcCtx: svcCtx,
	}
}

//  向手机号发送验证码
func (s *MessageServer) SentMessageCode(ctx context.Context, in *message.SentMessageCodeRequest) (*message.SentMessageCodeResponse, error) {
	l := logic.NewSentMessageCodeLogic(ctx, s.svcCtx)
	return l.SentMessageCode(in)
}

//  校验验证码
func (s *MessageServer) CheckMessageCode(ctx context.Context, in *message.CheckMessageCodeRequest) (*message.CheckMessageCodeResponse, error) {
	l := logic.NewCheckMessageCodeLogic(ctx, s.svcCtx)
	return l.CheckMessageCode(in)
}