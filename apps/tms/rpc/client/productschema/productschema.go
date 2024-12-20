// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package productschema

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProductCategoryReq                        = pb.CreateProductCategoryReq
	CreateProductCategoryResp                       = pb.CreateProductCategoryResp
	CreateProductEndpointReq                        = pb.CreateProductEndpointReq
	CreateProductEndpointResp                       = pb.CreateProductEndpointResp
	CreateProductReq                                = pb.CreateProductReq
	CreateProductResp                               = pb.CreateProductResp
	CreateProductSchemaModuleReq                    = pb.CreateProductSchemaModuleReq
	CreateProductSchemaModuleResp                   = pb.CreateProductSchemaModuleResp
	DeleteProductCategoryReq                        = pb.DeleteProductCategoryReq
	DeleteProductCategoryResp                       = pb.DeleteProductCategoryResp
	DeleteProductEndpointReq                        = pb.DeleteProductEndpointReq
	DeleteProductEndpointResp                       = pb.DeleteProductEndpointResp
	DeleteProductReq                                = pb.DeleteProductReq
	DeleteProductResp                               = pb.DeleteProductResp
	DeleteProductSchemaModuleReq                    = pb.DeleteProductSchemaModuleReq
	DeleteProductSchemaModuleResp                   = pb.DeleteProductSchemaModuleResp
	GetProductCategoryReq                           = pb.GetProductCategoryReq
	GetProductCategoryResp                          = pb.GetProductCategoryResp
	GetProductCategoryTreeReq                       = pb.GetProductCategoryTreeReq
	GetProductCategoryTreeResp                      = pb.GetProductCategoryTreeResp
	GetProductCategoryTreeResp_ProductCategory      = pb.GetProductCategoryTreeResp_ProductCategory
	GetProductEndpointReq                           = pb.GetProductEndpointReq
	GetProductEndpointResp                          = pb.GetProductEndpointResp
	GetProductReq                                   = pb.GetProductReq
	GetProductResp                                  = pb.GetProductResp
	GetProductSchemaModuleReq                       = pb.GetProductSchemaModuleReq
	GetProductSchemaModuleResp                      = pb.GetProductSchemaModuleResp
	GetProductSchemaVersionReq                      = pb.GetProductSchemaVersionReq
	GetProductSchemaVersionResp                     = pb.GetProductSchemaVersionResp
	GetProductSchemaVersionWithReq                  = pb.GetProductSchemaVersionWithReq
	GetProductSchemaVersionWithResp                 = pb.GetProductSchemaVersionWithResp
	ListProductEndpointReq                          = pb.ListProductEndpointReq
	ListProductEndpointResp                         = pb.ListProductEndpointResp
	ListProductEndpointResp_ProductEndpoint         = pb.ListProductEndpointResp_ProductEndpoint
	ListProductReq                                  = pb.ListProductReq
	ListProductResp                                 = pb.ListProductResp
	ListProductResp_Product                         = pb.ListProductResp_Product
	ListProductSchemaModuleReq                      = pb.ListProductSchemaModuleReq
	ListProductSchemaModuleResp                     = pb.ListProductSchemaModuleResp
	ListProductSchemaModuleResp_ProductSchemaModule = pb.ListProductSchemaModuleResp_ProductSchemaModule
	ListProductSchemaReq                            = pb.ListProductSchemaReq
	ListProductSchemaResp                           = pb.ListProductSchemaResp
	ListProductSchemaResp_ProductSchema             = pb.ListProductSchemaResp_ProductSchema
	UpdateProductCategoryReq                        = pb.UpdateProductCategoryReq
	UpdateProductCategoryResp                       = pb.UpdateProductCategoryResp
	UpdateProductEndpointReq                        = pb.UpdateProductEndpointReq
	UpdateProductEndpointResp                       = pb.UpdateProductEndpointResp
	UpdateProductReq                                = pb.UpdateProductReq
	UpdateProductResp                               = pb.UpdateProductResp
	UpdateProductSchemaModuleReq                    = pb.UpdateProductSchemaModuleReq
	UpdateProductSchemaModuleResp                   = pb.UpdateProductSchemaModuleResp
	UpdateProductStatusReq                          = pb.UpdateProductStatusReq
	UpdateProductStatusResp                         = pb.UpdateProductStatusResp

	ProductSchema interface {
		// 产品物模型列表
		ListProductSchema(ctx context.Context, in *ListProductSchemaReq, opts ...grpc.CallOption) (*ListProductSchemaResp, error)
	}

	defaultProductSchema struct {
		cli zrpc.Client
	}
)

func NewProductSchema(cli zrpc.Client) ProductSchema {
	return &defaultProductSchema{
		cli: cli,
	}
}

// 产品物模型列表
func (m *defaultProductSchema) ListProductSchema(ctx context.Context, in *ListProductSchemaReq, opts ...grpc.CallOption) (*ListProductSchemaResp, error) {
	client := pb.NewProductSchemaClient(m.cli.Conn())
	return client.ListProductSchema(ctx, in, opts...)
}