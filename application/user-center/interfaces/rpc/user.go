package main

import (
	"Ai-HireSphere/application/user-center/interfaces/rpc/types"
	"Ai-HireSphere/common/interceptors"
	"Ai-HireSphere/common/zlog"
	"flag"
	"fmt"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/config"
	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/server"
	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero --home=../../../../template

var configFile = flag.String("f", "./etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		types.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	defer s.Stop()
	//注册自定义日志
	zlog.InitLogger(c.ServiceConf)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
