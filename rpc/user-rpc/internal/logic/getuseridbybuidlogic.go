package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
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
	//user_info, err := user_service.FindByBuid(l.ctx, in.Buid)

	user_info, err := l.svcCtx.BiliUsersModel.FindOneByBuid(l.ctx, in.Buid)
	if err != nil {
		return &user.GetUserIDByBuidResponse{
			Code:   400,
			Msg:    "Get user_id failed",
			UserId: "",
		}, err
	}
	return &user.GetUserIDByBuidResponse{
		Code:   200,
		Msg:    "Success.",
		UserId: user_info.UserId.String,
	}, nil
}
