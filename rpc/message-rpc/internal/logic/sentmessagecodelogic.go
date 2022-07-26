package logic

import (
	"context"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"
	msg_service "github.com/leaper-one/2SOMEone/service/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type SentMessageCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSentMessageCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SentMessageCodeLogic {
	return &SentMessageCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  向手机号发送验证码
func (l *SentMessageCodeLogic) SentMessageCode(in *message.SentMessageCodeRequest) (*message.SentMessageCodeResponse, error) {
	code, err := msg_service.SendPhoneCode(l.ctx, in.Phone)

	if err != nil {
		return &message.SentMessageCodeResponse{
			Code:  500,
			Msg:   "发送验证码失败",
			MsgId: 0,
		}, err
	}

	ret, err := l.svcCtx.Model.InsertCode(context.Background(), in.Phone, code)

	if err != nil {
		return &message.SentMessageCodeResponse{
			Code:  500,
			Msg:   "存储验证码失败",
			MsgId: 0,
		}, err
	}
	msgId, _ := ret.LastInsertId()

	return &message.SentMessageCodeResponse{
		Code:  200,
		Msg:   "success",
		MsgId: uint64(msgId),
	}, nil
}
