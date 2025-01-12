// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package server

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/logic/productcategory"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
)

type ProductCategoryServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductCategoryServer
}

func NewProductCategoryServer(svcCtx *svc.ServiceContext) *ProductCategoryServer {
	return &ProductCategoryServer{
		svcCtx: svcCtx,
	}
}

// 创建产品分类
func (s *ProductCategoryServer) CreateProductCategory(ctx context.Context, in *pb.CreateProductCategoryReq) (*pb.CreateProductCategoryResp, error) {
	l := productcategorylogic.NewCreateProductCategoryLogic(ctx, s.svcCtx)
	return l.CreateProductCategory(in)
}

// 更新产品分类
func (s *ProductCategoryServer) UpdateProductCategory(ctx context.Context, in *pb.UpdateProductCategoryReq) (*pb.UpdateProductCategoryResp, error) {
	l := productcategorylogic.NewUpdateProductCategoryLogic(ctx, s.svcCtx)
	return l.UpdateProductCategory(in)
}

// 删除产品分类
func (s *ProductCategoryServer) DeleteProductCategory(ctx context.Context, in *pb.DeleteProductCategoryReq) (*pb.DeleteProductCategoryResp, error) {
	l := productcategorylogic.NewDeleteProductCategoryLogic(ctx, s.svcCtx)
	return l.DeleteProductCategory(in)
}

// 查询产品分类
func (s *ProductCategoryServer) GetProductCategory(ctx context.Context, in *pb.GetProductCategoryReq) (*pb.GetProductCategoryResp, error) {
	l := productcategorylogic.NewGetProductCategoryLogic(ctx, s.svcCtx)
	return l.GetProductCategory(in)
}

// 查询产品分类树
func (s *ProductCategoryServer) GetProductCategoryTree(ctx context.Context, in *pb.GetProductCategoryTreeReq) (*pb.GetProductCategoryTreeResp, error) {
	l := productcategorylogic.NewGetProductCategoryTreeLogic(ctx, s.svcCtx)
	return l.GetProductCategoryTree(in)
}
