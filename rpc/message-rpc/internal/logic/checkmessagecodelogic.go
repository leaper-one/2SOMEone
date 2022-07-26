package logic

import (
	"context"
	"time"

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

	code, phone, timestamp, err := l.svcCtx.Model.FindCode(context.Background(), int64(in.MsgId))

	if err != nil {
		return &message.CheckMessageCodeResponse{
			Code:    500,
			Msg:     "校验验证码失败",
			IsMatch: false,
		}, err
	}

	// 5分钟内有效
	if time.Now().Unix()-timestamp > 300 {
		return &message.CheckMessageCodeResponse{
			Code:    500,
			Msg:     "验证码已过期",
			IsMatch: false,
		}, nil
	}

	if phone != in.Phone {
		return &message.CheckMessageCodeResponse{
			Code:    200,
			Msg:     "手机号不匹配",
			IsMatch: false,
		}, nil
	}

	if code != in.Code {
		return &message.CheckMessageCodeResponse{
			Code:    200,
			Msg:     "验证码不匹配",
			IsMatch: false,
		}, nil
	}

	return &message.CheckMessageCodeResponse{
		Code:    200,
		Msg:     "success",
		IsMatch: true,
	}, nil
}
