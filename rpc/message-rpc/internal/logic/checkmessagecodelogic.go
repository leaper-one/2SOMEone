package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMessageCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckMessageCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMessageCodeLogic {
	return &CheckMessageCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  校验验证码
func (l *CheckMessageCodeLogic) CheckMessageCode(in *message.CheckMessageCodeRequest) (*message.CheckMessageCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &message.CheckMessageCodeResponse{}, nil
}
