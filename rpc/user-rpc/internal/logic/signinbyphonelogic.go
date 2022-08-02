package logic

import (
	"context"
	"github.com/leaper-one/2SOMEone/util"
	"time"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
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

	var token string

	userData, err := l.svcCtx.BasicUsersModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return &user.SignInByPhoneResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	if userData == nil {
		return &user.SignInByPhoneResponse{
			Code: 400,
			Msg:  "该手机号未注册",
		}, nil
	}

	if util.CheckPasswordHash(in.Password, userData.Password.String) {
		token, err = util.GenerateToken(userData.UserId.String, userData.Phone.String, l.svcCtx.Config.JwtAuth.AccessSecret, time.Duration(l.svcCtx.Config.JwtAuth.AccessExpire)*time.Second)
		if err != nil {
			return &user.SignInByPhoneResponse{
				Code: 500,
				Msg:  err.Error(),
			}, nil
		}
	} else {
		return &user.SignInByPhoneResponse{
			Code: 400,
			Msg:  "密码错误",
		}, nil
	}

	return &user.SignInByPhoneResponse{
		Code:  200,
		Msg:   "登录成功",
		Token: token,
	}, nil
}
