package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	user_service "github.com/leaper-one/2SOMEone/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpByPhoneLogic {
	return &SignUpByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过手机号注册，需要验证码
func (l *SignUpByPhoneLogic) SignUpByPhone(in *user.SignUpByPhoneRequest) (*user.SignUpByPhoneResponse, error) {
	// 调用 message-rpc 验证短信验证码
	res, err := l.svcCtx.Message.CheckMessageCode(l.ctx, &message.CheckMessageCodeRequest{
		Phone: in.Phone,
		Code:  in.Code,
		MsgId: in.MsgId,
	})
	if err != nil {
		return &user.SignUpByPhoneResponse{
			Code: 500,
			Msg:  "Falt.",
		}, nil
	}
	if !res.IsMatch {
		return &user.SignUpByPhoneResponse{
			Code: 400,
			Msg:  "验证码错误",
		}, nil

	}
	// 注册用户
	err = user_service.SignUpByPhone(l.ctx, in.Phone, in.Password)
	if err != nil {
		return &user.SignUpByPhoneResponse{
			Code: 500,
			Msg:  "注册失败",
		}, err
	}
	return &user.SignUpByPhoneResponse{
		Code: 200,
		Msg:  "注册成功",
	}, nil
}
