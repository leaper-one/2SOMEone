package svc

import (
	"github.com/leaper-one/2SOMEone/model/message"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  message.MessagesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  message.NewMessagesModel(sqlx.NewMysql(c.DataSource)),
	}
}
