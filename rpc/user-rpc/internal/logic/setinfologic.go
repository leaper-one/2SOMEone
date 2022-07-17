package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInfoLogic {
	return &SetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  jwt needed in metadata
func (l *SetInfoLogic) SetInfo(in *user.SetInfoRequest) (*user.SetInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SetInfoResponse{}, nil
}
