package accounts

import (
	"context"

	"github.com/leaper-one/2SOMEone/api/user-api/internal/svc"
	"github.com/leaper-one/2SOMEone/api/user-api/internal/types"

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

func (l *SentPhoneCodeLogic) SentPhoneCode() (resp *types.SentPhoneCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
