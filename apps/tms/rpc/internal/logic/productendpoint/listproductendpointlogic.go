package productendpointlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductEndpointLogic {
	return &ListProductEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListProductEndpoint 查看产品通信地址列表
func (l *ListProductEndpointLogic) ListProductEndpoint(in *pb.ListProductEndpointReq) (*pb.ListProductEndpointResp, error) {
	pes, err := l.svcCtx.ProductEndpointModel.List(l.ctx, in.ProductId)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var data []*pb.ListProductEndpointResp_ProductEndpoint
	_ = copier.CopyWithOption(&data, pes, xcopier.CopierOption)

	return &pb.ListProductEndpointResp{
		Data: data,
	}, nil
}
