package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/user/api/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/user/api/internal/types"

	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	res, err := l.svcCtx.User.SignUpByPhone(l.ctx, &user.SignUpByPhoneRequest{
		Phone:    req.Phone,
		Code:     req.Phone_code,
		Password: req.Password,
		MsgId:    req.Msg_id,
	})
	if err != nil {
		return &types.SignUpResp{}, err
	}
	return &types.SignUpResp{
		Code: res.Code,
		Msg:  res.Msg,
	},nil
}
