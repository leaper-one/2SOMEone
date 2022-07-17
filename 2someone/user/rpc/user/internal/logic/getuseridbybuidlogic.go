package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/internal/svc"
	"github.com/leaper-one/2SOMEone/2someone/user/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIDByBuidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserIDByBuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIDByBuidLogic {
	return &GetUserIDByBuidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据 buid 获取 user_id
func (l *GetUserIDByBuidLogic) GetUserIDByBuid(in *user.GetUserIDByBuidRequest) (*user.GetUserIDByBuidResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserIDByBuidResponse{}, nil
}
