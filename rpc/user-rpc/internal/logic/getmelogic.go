package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	user_service "github.com/leaper-one/2SOMEone/service/user"

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
	user_info, err := user_service.GetMe(l.ctx, in.UserId)
	if err != nil {
		return &user.GetMeResponse{
			Code:	400,
			Msg:	"Get user info failed",
			User: 	nil,
		}, err
	}

	return &user.GetMeResponse{
		Code:	200,
		Msg:	"Success.",
		User: 	&user.BasicUser{
			UserId: user_info.UserID,
			UserInfo: &user.UserInfo{
				Name: user_info.Name,
				Phone: user_info.Phone,
				Avatar: user_info.Avatar,
				Email: user_info.Email,
			},
		},
	}, nil
}
