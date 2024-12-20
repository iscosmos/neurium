// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tms.proto

package server

import (
	"context"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/logic/productschema"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
)

type ProductSchemaServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductSchemaServer
}

func NewProductSchemaServer(svcCtx *svc.ServiceContext) *ProductSchemaServer {
	return &ProductSchemaServer{
		svcCtx: svcCtx,
	}
}

// 产品物模型列表
func (s *ProductSchemaServer) ListProductSchema(ctx context.Context, in *pb.ListProductSchemaReq) (*pb.ListProductSchemaResp, error) {
	l := productschemalogic.NewListProductSchemaLogic(ctx, s.svcCtx)
	return l.ListProductSchema(in)
}
