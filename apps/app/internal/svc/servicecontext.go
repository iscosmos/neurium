package svc

import (
	"github.com/iscosmos/neurium/apps/app/internal/config"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschema"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemaversion"
	"github.com/iscosmos/neurium/pkg/interceptor"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                  config.Config
	ProductRpc              product.Product
	ProductCategoryRpc      productcategory.ProductCategory
	ProductEndpointRpc      productendpoint.ProductEndpoint
	ProductSchemaRpc        productschema.ProductSchema
	ProductSchemaVersionRpc productschemaversion.ProductSchemaVersion
	ProductSchemaModuleRpc  productschemamodule.ProductSchemaModule
}

func NewServiceContext(c config.Config) *ServiceContext {
	tmsRpcClient := zrpc.MustNewClient(c.TmsRpc, zrpc.WithUnaryClientInterceptor(interceptor.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:                  c,
		ProductRpc:              product.NewProduct(tmsRpcClient),
		ProductCategoryRpc:      productcategory.NewProductCategory(tmsRpcClient),
		ProductEndpointRpc:      productendpoint.NewProductEndpoint(tmsRpcClient),
		ProductSchemaRpc:        productschema.NewProductSchema(tmsRpcClient),
		ProductSchemaVersionRpc: productschemaversion.NewProductSchemaVersion(tmsRpcClient),
		ProductSchemaModuleRpc:  productschemamodule.NewProductSchemaModule(tmsRpcClient),
	}
}
