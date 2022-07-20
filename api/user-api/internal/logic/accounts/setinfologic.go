package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"

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
	logx.Info("userId: %d", l.ctx.Value("user_id"))
	// fmt.Printf("l.ctx.Value(\"user_id\").(string): %v\n", l.ctx.Value("user_id").(string))
	_, err = l.svcCtx.User.SetInfo(l.ctx, &user.SetInfoRequest{
		Name:   req.Name,
		Avatar: req.Avatar,
		Buid:   req.Buid,
		UserId: l.ctx.Value("user_id").(string),
	})
	if err != nil {
		return &types.SetInfoResp{
			Code: 400,
			Msg:  "set info failed",
		}, err
	}
	return &types.SetInfoResp{
		Code: 200,
		Msg:  "success",
	}, nil
}
