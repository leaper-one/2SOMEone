package app

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIdByBuidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserIdByBuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIdByBuidLogic {
	return &GetUserIdByBuidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserIdByBuidLogic) GetUserIdByBuid(req *types.GetUserIdByBuidReq) (resp *types.GetUserIdByBuidResp, err error) {
	res, err := l.svcCtx.User.GetUserIDByBuid(l.ctx, &user.GetUserIDByBuidRequest{
		Buid: req.Buid,
	})
	if err != nil {
		return &types.GetUserIdByBuidResp{}, err
	}
	return &types.GetUserIdByBuidResp{
		Code:   res.Code,
		Msg:    res.Msg,
		UserId: res.UserId,
	}, nil
}
