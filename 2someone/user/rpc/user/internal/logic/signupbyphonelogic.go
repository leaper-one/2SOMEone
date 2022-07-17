package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/types/user"

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
	// todo: add your logic here and delete this line

	return &user.SignUpByPhoneResponse{}, nil
}
