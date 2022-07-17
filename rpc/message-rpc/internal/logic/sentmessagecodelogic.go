package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/gen/go/message"

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
	// todo: add your logic here and delete this line

	return &message.SentMessageCodeResponse{}, nil
}
