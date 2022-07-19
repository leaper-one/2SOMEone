package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInfoLogic {
	return &SetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetInfoLogic) SetInfo(req *types.SetInfoReq) (resp *types.SetInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
