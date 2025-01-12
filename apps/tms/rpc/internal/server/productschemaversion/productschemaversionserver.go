// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package server

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/logic/productschemaversion"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
)

type ProductSchemaVersionServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductSchemaVersionServer
}

func NewProductSchemaVersionServer(svcCtx *svc.ServiceContext) *ProductSchemaVersionServer {
	return &ProductSchemaVersionServer{
		svcCtx: svcCtx,
	}
}

// 查看物模型版本
func (s *ProductSchemaVersionServer) GetProductSchemaVersion(ctx context.Context, in *pb.GetProductSchemaVersionReq) (*pb.GetProductSchemaVersionResp, error) {
	l := productschemaversionlogic.NewGetProductSchemaVersionLogic(ctx, s.svcCtx)
	return l.GetProductSchemaVersion(in)
}

// 根据产品和状态查询物模型版本
func (s *ProductSchemaVersionServer) GetProductSchemaVersionWith(ctx context.Context, in *pb.GetProductSchemaVersionWithReq) (*pb.GetProductSchemaVersionWithResp, error) {
	l := productschemaversionlogic.NewGetProductSchemaVersionWithLogic(ctx, s.svcCtx)
	return l.GetProductSchemaVersionWith(in)
}
