package logic

import (
	"context"
	"database/sql"

	"github.com/leaper-one/2SOMEone/model/basic_user"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/types/user"
	user_service "github.com/leaper-one/2SOMEone/service/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpByPhoneLogic {
	return &SignUpByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  通过手机号注册，需要验证码
func (l *SignUpByPhoneLogic) SignUpByPhone(in *user.SignUpByPhoneRequest) (*user.SignUpByPhoneResponse, error) {
	// 调用 message-rpc 验证短信验证码
	res, err := l.svcCtx.Message.CheckMessageCode(l.ctx, &message.CheckMessageCodeRequest{
		Phone: in.Phone,
		Code:  in.Code,
		MsgId: in.MsgId,
	})
	if err != nil {
		return &user.SignUpByPhoneResponse{
			Code: 500,
			Msg:  "Falt.",
		}, nil
	}
	if !res.IsMatch {
		return &user.SignUpByPhoneResponse{
			Code: 400,
			Msg:  "验证码错误",
		}, nil

	}
	// 注册用户
	// err = user_service.SignUpByPhone(l.ctx, in.Phone, in.Password)
	// if err != nil {
	// 	return &user.SignUpByPhoneResponse{
	// 		Code: 500,
	// 		Msg:  "注册失败",
	// 	}, err
	// }
	// return &user.SignUpByPhoneResponse{
	// 	Code: 200,
	// 	Msg:  "注册成功",
	// }, nil

	// 查询用户是否存在
	_, err =l.svcCtx.Model.FindOneByPhone(l.ctx, in.Phone)
	if err == basic_user.ErrNotFound {
		return &user.SignUpByPhoneResponse{
			Code: 400,
			Msg:  "用户已存在",
		}, nil
	}

	// 创建用户
	user_id, password, err := user_service.CreateUser(l.ctx, in.Password)
	if err != nil {
		return &user.SignUpByPhoneResponse{
			Code: 500,
			Msg:  "创建用户失败",
		}, err
	}
	// 插入数据库
	_, err = l.svcCtx.Model.Insert(l.ctx, &basic_user.BasicUsers{
		UserId:   sql.NullString{String: user_id, Valid: true},
		Phone:    sql.NullString{String: in.Phone, Valid: true},
		Password: sql.NullString{String: password, Valid: true},
		Name:     sql.NullString{String: in.Phone, Valid: true},
	})

	if err != nil {
		return &user.SignUpByPhoneResponse{
			Code: 500,
			Msg:  "存储用户失败",
		}, err
	}

	return &user.SignUpByPhoneResponse{
		Code: 200,
		Msg:  "注册成功",
	}, nil

}
