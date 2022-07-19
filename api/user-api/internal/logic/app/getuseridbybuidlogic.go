package app

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
