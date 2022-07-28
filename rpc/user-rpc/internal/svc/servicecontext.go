package svc

import (
	"github.com/leaper-one/2SOMEone/model/basic_user"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/message"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Model   basic_user.BasicUsersModel
	Message message.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Model:   basic_user.NewBasicUsersModel(sqlx.NewMysql(c.DataSource)),
		Message: message.NewMessage(zrpc.MustNewClient(c.Message)),
	}
}
