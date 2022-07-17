package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/message/rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/message/rpc/types/message"
	msg "github.com/leaper-one/2SOMEone/2someone/message/service/message"

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
	is_match, err := msg.CheckPhoneCode(l.ctx, in.Phone, in.Code, uint(in.MsgId))
	if err != nil {
		return &message.CheckMessageCodeResponse{
			Code:    500,
			Msg:     "failed",
			IsMatch: false,
		}, err
	}

	return &message.CheckMessageCodeResponse{
		Code:    200,
		Msg:     "success",
		IsMatch: is_match,
	}, nil
}
