// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package productschemaversion

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

	ProductSchemaVersion interface {
		// 查看物模型版本
		GetProductSchemaVersion(ctx context.Context, in *GetProductSchemaVersionReq, opts ...grpc.CallOption) (*GetProductSchemaVersionResp, error)
		// 根据产品和状态查询物模型版本
		GetProductSchemaVersionWith(ctx context.Context, in *GetProductSchemaVersionWithReq, opts ...grpc.CallOption) (*GetProductSchemaVersionWithResp, error)
	}

	defaultProductSchemaVersion struct {
		cli zrpc.Client
	}
)

func NewProductSchemaVersion(cli zrpc.Client) ProductSchemaVersion {
	return &defaultProductSchemaVersion{
		cli: cli,
	}
}

// 查看物模型版本
func (m *defaultProductSchemaVersion) GetProductSchemaVersion(ctx context.Context, in *GetProductSchemaVersionReq, opts ...grpc.CallOption) (*GetProductSchemaVersionResp, error) {
	client := pb.NewProductSchemaVersionClient(m.cli.Conn())
	return client.GetProductSchemaVersion(ctx, in, opts...)
}

// 根据产品和状态查询物模型版本
func (m *defaultProductSchemaVersion) GetProductSchemaVersionWith(ctx context.Context, in *GetProductSchemaVersionWithReq, opts ...grpc.CallOption) (*GetProductSchemaVersionWithResp, error) {
	client := pb.NewProductSchemaVersionClient(m.cli.Conn())
	return client.GetProductSchemaVersionWith(ctx, in, opts...)
}
