package logic

import (
	"context"

	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	user_service "github.com/leaper-one/2SOMEone/service/user"

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
	// logx.Info("rpc user-rpc userId: %d", l.ctx.Value("user_id"))
	// user_id := l.ctx.Value("user_id").(string)
	err := user_service.SetInfo(l.ctx, in.UserId, in.Name, in.Avatar, in.Buid)
	if err != nil {
		return &user.SetInfoResponse{
			Code: 400,
			Msg: "set info failed",
		}, err
	}
	return &user.SetInfoResponse{
		Code: 200,
		Msg: "success",
	}, nil
}
