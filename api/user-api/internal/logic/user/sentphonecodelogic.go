package user

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"
	"github.com/zeromicro/go-zero/core/logx"
)

type SentPhoneCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSentPhoneCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SentPhoneCodeLogic {
	return &SentPhoneCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SentPhoneCodeLogic) SentPhoneCode(req *types.SentPhoneCodeReq) (resp *types.SentPhoneCodeResp, err error) {
	res, err := l.svcCtx.Message.SentMessageCode(l.ctx, &message.SentMessageCodeRequest{
		Phone: req.Phone,
	})
	if err != nil {
		return &types.SentPhoneCodeResp{}, err
	}
	return &types.SentPhoneCodeResp{
		Code: res.Code,
		Msg:  res.Msg,
		Msg_id: res.MsgId,
	}, nil

}
