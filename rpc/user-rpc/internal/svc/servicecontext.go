package svc

import (
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/message"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Message message.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Message: message.NewMessage(zrpc.MustNewClient(c.Message)),
	}
}
