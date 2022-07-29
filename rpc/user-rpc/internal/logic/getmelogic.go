package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Get current user infomation by metadata with auth token
func (l *GetMeLogic) GetMe(in *user.GetMeRequest) (*user.GetMeResponse, error) {
	user_info, err := l.svcCtx.BasicUsersModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return &user.GetMeResponse{
			Code: 400,
			Msg:  "Get user info failed",
			User: nil,
		}, err
	}

	return &user.GetMeResponse{
		Code: 200,
		Msg:  "Success.",
		User: &user.BasicUser{
			UserId: user_info.UserId.String,
			UserInfo: &user.UserInfo{
				Name:   user_info.Name.String,
				Phone:  user_info.Phone.String,
				Avatar: user_info.Avatar.String,
				Email:  user_info.Email.String,
			},
		},
	}, nil
}
