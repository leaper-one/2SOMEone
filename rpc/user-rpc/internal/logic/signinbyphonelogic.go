package logic

import (
	"context"
	"time"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	user_service "github.com/leaper-one/2SOMEone/service/user"

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
	// 调用 user service 验证用户名和密码
	token, err := user_service.Auth(l.ctx, in.Phone, in.Password, l.svcCtx.Config.JwtAuth.AccessSecret, time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire) * time.Second)
	if err != nil {
		return &user.SignInByPhoneResponse{
			Code: 400,
			Msg:  "登录失败",
		}, err
	}
	return &user.SignInByPhoneResponse{
		Code:  200,
		Msg:   "登录成功",
		Token: token,
	}, nil
}
