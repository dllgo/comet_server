package main

import (
	"flag"
	"fmt"

	"github.com/dllgo/comet"
	"github.com/dllgo/comet_server/comet_server"
	"github.com/dllgo/comet_server/internal/config"
	"github.com/dllgo/comet_server/internal/server"
	"github.com/dllgo/comet_server/internal/svc"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/cometserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCometServerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comet_server.RegisterCometServerServer(grpcServer, srv)
	})
	defer s.Stop()


	//comet server  
	go comet.NewComet().Serve(server.NewMessageHandler(ctx),c.CometPort)
	//
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
