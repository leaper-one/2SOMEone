package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMeLogic) GetMe() (resp *types.MeResp, err error) {
	logx.Info("userId: %d", l.ctx.Value("user_id`"))
	user_id := l.ctx.Value("user_id").(string)
	res, err := l.svcCtx.User.GetMe(l.ctx, &user.GetMeRequest{
		UserId: user_id,
	})
	if err != nil {
		return &types.MeResp{}, err
	}
	return &types.MeResp{
		Code: res.Code,
		Msg: res.Msg,
		UserId: res.User.UserId,
		Name: res.User.UserInfo.Name,
		Phone: res.User.UserInfo.Phone,
		Avatar: res.User.UserInfo.Avatar,
		Email: res.User.UserInfo.Email,
	}, nil
}
