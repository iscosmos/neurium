// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package productendpoint

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

	ProductEndpoint interface {
		// 创建产品通信地址
		CreateProductEndpoint(ctx context.Context, in *CreateProductEndpointReq, opts ...grpc.CallOption) (*CreateProductEndpointResp, error)
		// 更新产品通信地址
		UpdateProductEndpoint(ctx context.Context, in *UpdateProductEndpointReq, opts ...grpc.CallOption) (*UpdateProductEndpointResp, error)
		// 删除产品通信地址
		DeleteProductEndpoint(ctx context.Context, in *DeleteProductEndpointReq, opts ...grpc.CallOption) (*DeleteProductEndpointResp, error)
		// 查看产品通信地址
		GetProductEndpoint(ctx context.Context, in *GetProductEndpointReq, opts ...grpc.CallOption) (*GetProductEndpointResp, error)
		// 查看产品通信地址列表
		ListProductEndpoint(ctx context.Context, in *ListProductEndpointReq, opts ...grpc.CallOption) (*ListProductEndpointResp, error)
	}

	defaultProductEndpoint struct {
		cli zrpc.Client
	}
)

func NewProductEndpoint(cli zrpc.Client) ProductEndpoint {
	return &defaultProductEndpoint{
		cli: cli,
	}
}

// 创建产品通信地址
func (m *defaultProductEndpoint) CreateProductEndpoint(ctx context.Context, in *CreateProductEndpointReq, opts ...grpc.CallOption) (*CreateProductEndpointResp, error) {
	client := pb.NewProductEndpointClient(m.cli.Conn())
	return client.CreateProductEndpoint(ctx, in, opts...)
}

// 更新产品通信地址
func (m *defaultProductEndpoint) UpdateProductEndpoint(ctx context.Context, in *UpdateProductEndpointReq, opts ...grpc.CallOption) (*UpdateProductEndpointResp, error) {
	client := pb.NewProductEndpointClient(m.cli.Conn())
	return client.UpdateProductEndpoint(ctx, in, opts...)
}

// 删除产品通信地址
func (m *defaultProductEndpoint) DeleteProductEndpoint(ctx context.Context, in *DeleteProductEndpointReq, opts ...grpc.CallOption) (*DeleteProductEndpointResp, error) {
	client := pb.NewProductEndpointClient(m.cli.Conn())
	return client.DeleteProductEndpoint(ctx, in, opts...)
}

// 查看产品通信地址
func (m *defaultProductEndpoint) GetProductEndpoint(ctx context.Context, in *GetProductEndpointReq, opts ...grpc.CallOption) (*GetProductEndpointResp, error) {
	client := pb.NewProductEndpointClient(m.cli.Conn())
	return client.GetProductEndpoint(ctx, in, opts...)
}

// 查看产品通信地址列表
func (m *defaultProductEndpoint) ListProductEndpoint(ctx context.Context, in *ListProductEndpointReq, opts ...grpc.CallOption) (*ListProductEndpointResp, error) {
	client := pb.NewProductEndpointClient(m.cli.Conn())
	return client.ListProductEndpoint(ctx, in, opts...)
}
