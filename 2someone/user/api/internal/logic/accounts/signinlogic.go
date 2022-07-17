package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/user/api/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/user/api/internal/types"
	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (*types.SignInResp, error) {
	resp, err := l.svcCtx.User.SignInByPhone(l.ctx, &user.SignInByPhoneRequest{
		Phone: req.Phone,
		Password:  req.Password,
	})
	if err != nil {
		return &types.SignInResp{}, err
	}
	return &types.SignInResp{
		Token: resp.Token,
	}, nil
}
