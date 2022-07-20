package main

import (
	"flag"
	"fmt"

	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/config"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/server"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/internal/svc"
	"github.com/leaper-one/2SOMEone/rpc/message-rpc/types/message"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewMessageServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		message.RegisterMessageServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
