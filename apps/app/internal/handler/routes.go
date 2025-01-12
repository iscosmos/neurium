// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	tmsproduct "github.com/iscosmos/neurium/apps/app/internal/handler/tms/product"
	tmsproductcategory "github.com/iscosmos/neurium/apps/app/internal/handler/tms/productcategory"
	tmsproductendpoint "github.com/iscosmos/neurium/apps/app/internal/handler/tms/productendpoint"
	tmsproductschema "github.com/iscosmos/neurium/apps/app/internal/handler/tms/productschema"
	tmsproductschemamodule "github.com/iscosmos/neurium/apps/app/internal/handler/tms/productschemamodule"
	tmsproductschemaversion "github.com/iscosmos/neurium/apps/app/internal/handler/tms/productschemaversion"
	"github.com/iscosmos/neurium/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建产品
				Method:  http.MethodPost,
				Path:    "/product",
				Handler: tmsproduct.CreateProductHandler(serverCtx),
			},
			{
				// 更新产品
				Method:  http.MethodPut,
				Path:    "/product",
				Handler: tmsproduct.UpdateProductHandler(serverCtx),
			},
			{
				// 删除产品
				Method:  http.MethodDelete,
				Path:    "/product",
				Handler: tmsproduct.DeleteProductHandler(serverCtx),
			},
			{
				// 产品列表
				Method:  http.MethodGet,
				Path:    "/product",
				Handler: tmsproduct.ListProductHandler(serverCtx),
			},
			{
				// 产品详情
				Method:  http.MethodGet,
				Path:    "/product/:id",
				Handler: tmsproduct.GetProductHandler(serverCtx),
			},
			{
				// 修改发布状态
				Method:  http.MethodPost,
				Path:    "/product/status",
				Handler: tmsproduct.UpdateProductStatusHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建产品分类
				Method:  http.MethodPost,
				Path:    "/productcategory",
				Handler: tmsproductcategory.CreateProductCategoryHandler(serverCtx),
			},
			{
				// 更新产品分类
				Method:  http.MethodPut,
				Path:    "/productcategory",
				Handler: tmsproductcategory.UpdateProductCategoryHandler(serverCtx),
			},
			{
				// 删除产品分类
				Method:  http.MethodDelete,
				Path:    "/productcategory",
				Handler: tmsproductcategory.DeleteProductCategoryHandler(serverCtx),
			},
			{
				// 查询产品分类
				Method:  http.MethodGet,
				Path:    "/productcategory/:id",
				Handler: tmsproductcategory.GetProductCategoryHandler(serverCtx),
			},
			{
				// 查询产品分类树
				Method:  http.MethodGet,
				Path:    "/productcategory/tree",
				Handler: tmsproductcategory.GetProductCategoryTreeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建产品通信地址
				Method:  http.MethodPost,
				Path:    "/productendpoint",
				Handler: tmsproductendpoint.CreateProductEndpointHandler(serverCtx),
			},
			{
				// 更新产品通信地址
				Method:  http.MethodPut,
				Path:    "/productendpoint",
				Handler: tmsproductendpoint.UpdateProductEndpointHandler(serverCtx),
			},
			{
				// 查询产品通信地址列表
				Method:  http.MethodGet,
				Path:    "/productendpoint",
				Handler: tmsproductendpoint.ListProductEndpointHandler(serverCtx),
			},
			{
				// 删除产品通信地址
				Method:  http.MethodDelete,
				Path:    "/productendpoint/:id",
				Handler: tmsproductendpoint.DeleteProductEndpointHandler(serverCtx),
			},
			{
				// 产品通信地址详情
				Method:  http.MethodGet,
				Path:    "/productendpoint/:id",
				Handler: tmsproductendpoint.GetProductEndpointHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 物模型列表
				Method:  http.MethodGet,
				Path:    "/productschema",
				Handler: tmsproductschema.ListProductSchemaHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建物模型模块
				Method:  http.MethodPost,
				Path:    "/productschemamodule",
				Handler: tmsproductschemamodule.CreateProductSchemaModuleHandler(serverCtx),
			},
			{
				// 更新物模型模块
				Method:  http.MethodPut,
				Path:    "/productschemamodule",
				Handler: tmsproductschemamodule.UpdateProductSchemaModuleHandler(serverCtx),
			},
			{
				// 删除物模型模块
				Method:  http.MethodDelete,
				Path:    "/productschemamodule",
				Handler: tmsproductschemamodule.DeleteProductSchemaModuleHandler(serverCtx),
			},
			{
				// 物模型模块列表
				Method:  http.MethodGet,
				Path:    "/productschemamodule",
				Handler: tmsproductschemamodule.ListProductSchemaModuleHandler(serverCtx),
			},
			{
				// 查询物模型模块
				Method:  http.MethodGet,
				Path:    "/productschemamodule/:id",
				Handler: tmsproductschemamodule.GetProductSchemaModuleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 根据产品和状态查询物模型版本
				Method:  http.MethodGet,
				Path:    "/productschemaversionwith",
				Handler: tmsproductschemaversion.GetProductSchemaVersionWithHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/tms"),
	)
}
