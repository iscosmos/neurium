package main

import (
	"flag"
	"fmt"
	"github.com/iscosmos/neurium/pkg/interceptor"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/config"
	productServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/product"
	productcategoryServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/productcategory"
	productEndpointServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/productendpoint"
	productSchemaServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/productschema"
	productSchemaModuleServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/productschemamodule"
	productSchemaVersionServer "github.com/iscosmos/neurium/apps/tms/rpc/internal/server/productschemaversion"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/tms.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterProductServer(grpcServer, productServer.NewProductServer(ctx))
		pb.RegisterProductCategoryServer(grpcServer, productcategoryServer.NewProductCategoryServer(ctx))
		pb.RegisterProductEndpointServer(grpcServer, productEndpointServer.NewProductEndpointServer(ctx))
		pb.RegisterProductSchemaVersionServer(grpcServer, productSchemaVersionServer.NewProductSchemaVersionServer(ctx))
		pb.RegisterProductSchemaModuleServer(grpcServer, productSchemaModuleServer.NewProductSchemaModuleServer(ctx))
		pb.RegisterProductSchemaServer(grpcServer, productSchemaServer.NewProductSchemaServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 添加错误处理自定义拦截器
	s.AddUnaryInterceptors(interceptor.ServerErrorInterceptor())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
