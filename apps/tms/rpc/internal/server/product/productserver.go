// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package server

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/logic/product"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

// 创建产品
func (s *ProductServer) CreateProduct(ctx context.Context, in *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	l := productlogic.NewCreateProductLogic(ctx, s.svcCtx)
	return l.CreateProduct(in)
}

// 更新产品
func (s *ProductServer) UpdateProduct(ctx context.Context, in *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	l := productlogic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

// 删除产品
func (s *ProductServer) DeleteProduct(ctx context.Context, in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	l := productlogic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}

// 产品详情
func (s *ProductServer) GetProduct(ctx context.Context, in *pb.GetProductReq) (*pb.GetProductResp, error) {
	l := productlogic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

// 产品列表
func (s *ProductServer) ListProduct(ctx context.Context, in *pb.ListProductReq) (*pb.ListProductResp, error) {
	l := productlogic.NewListProductLogic(ctx, s.svcCtx)
	return l.ListProduct(in)
}

// 修改发布状态
func (s *ProductServer) UpdateProductStatus(ctx context.Context, in *pb.UpdateProductStatusReq) (*pb.UpdateProductStatusResp, error) {
	l := productlogic.NewUpdateProductStatusLogic(ctx, s.svcCtx)
	return l.UpdateProductStatus(in)
}