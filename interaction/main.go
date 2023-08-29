package main

import (
	"flag"
	"fmt"
	"github.com/sunflower10086/TikTok/interaction/internal/config"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/db"
	"github.com/sunflower10086/TikTok/interaction/internal/server"
	"github.com/sunflower10086/TikTok/interaction/internal/svc"
	___interaction "github.com/sunflower10086/TikTok/interaction/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/interaction.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	config.LoadConfigFromYaml(*configFile)

	err := db.Init()
	if err != nil {
		fmt.Println("数据库连接失败!")
		return
	}

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		___interaction.RegisterInteractionServer(grpcServer, server.NewInteractionServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
