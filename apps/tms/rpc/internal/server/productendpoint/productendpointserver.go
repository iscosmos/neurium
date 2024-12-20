// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package server

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/logic/productendpoint"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
)

type ProductEndpointServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductEndpointServer
}

func NewProductEndpointServer(svcCtx *svc.ServiceContext) *ProductEndpointServer {
	return &ProductEndpointServer{
		svcCtx: svcCtx,
	}
}

// 创建产品通信地址
func (s *ProductEndpointServer) CreateProductEndpoint(ctx context.Context, in *pb.CreateProductEndpointReq) (*pb.CreateProductEndpointResp, error) {
	l := productendpointlogic.NewCreateProductEndpointLogic(ctx, s.svcCtx)
	return l.CreateProductEndpoint(in)
}

// 更新产品通信地址
func (s *ProductEndpointServer) UpdateProductEndpoint(ctx context.Context, in *pb.UpdateProductEndpointReq) (*pb.UpdateProductEndpointResp, error) {
	l := productendpointlogic.NewUpdateProductEndpointLogic(ctx, s.svcCtx)
	return l.UpdateProductEndpoint(in)
}

// 删除产品通信地址
func (s *ProductEndpointServer) DeleteProductEndpoint(ctx context.Context, in *pb.DeleteProductEndpointReq) (*pb.DeleteProductEndpointResp, error) {
	l := productendpointlogic.NewDeleteProductEndpointLogic(ctx, s.svcCtx)
	return l.DeleteProductEndpoint(in)
}

// 查看产品通信地址
func (s *ProductEndpointServer) GetProductEndpoint(ctx context.Context, in *pb.GetProductEndpointReq) (*pb.GetProductEndpointResp, error) {
	l := productendpointlogic.NewGetProductEndpointLogic(ctx, s.svcCtx)
	return l.GetProductEndpoint(in)
}

// 查看产品通信地址列表
func (s *ProductEndpointServer) ListProductEndpoint(ctx context.Context, in *pb.ListProductEndpointReq) (*pb.ListProductEndpointResp, error) {
	l := productendpointlogic.NewListProductEndpointLogic(ctx, s.svcCtx)
	return l.ListProductEndpoint(in)
}
