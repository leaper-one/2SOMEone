package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/gen/go/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInByPhoneLogic {
	return &SignInByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过手机号登录
func (l *SignInByPhoneLogic) SignInByPhone(in *user.SignInByPhoneRequest) (*user.SignInByPhoneResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SignInByPhoneResponse{}, nil
}
