package productendpoint

import (
	"context"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductEndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductEndpointLogic {
	return &ListProductEndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListProductEndpoint 查询产品通信地址列表
func (l *ListProductEndpointLogic) ListProductEndpoint(req *types.ListProductEndpointReq) (*types.ListProductEndpointResp, error) {
	out, err := l.svcCtx.ProductEndpointRpc.ListProductEndpoint(l.ctx, &productendpoint.ListProductEndpointReq{
		ProductId: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	var data []types.ProductEndpointItem
	_ = copier.CopyWithOption(&data, out.Data, xcopier.CopierOption)

	return &types.ListProductEndpointResp{Data: data}, nil
}
