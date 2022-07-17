package svc

import "github.com/leaper-one/2SOMEone/2someone/user/rpc/user/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
