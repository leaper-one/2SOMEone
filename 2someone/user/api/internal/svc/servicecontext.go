package svc

import (
	// "github.com/leaper-one/2SOMEone/2someone/message/rpc/message"
	"github.com/leaper-one/2SOMEone/2someone/user/api/internal/config"
	user "github.com/leaper-one/2SOMEone/2someone/user/rpc/user/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// Message message.Message
	User user.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   user.NewUserService(zrpc.MustNewClient(c.User)),
		// Message: message.NewMessage(zrpc.MustNewClient(c.User)),
	}
}
