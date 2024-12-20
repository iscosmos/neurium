package productendpoint

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductEndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductEndpointLogic {
	return &GetProductEndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductEndpoint 产品通信地址详情
func (l *GetProductEndpointLogic) GetProductEndpoint(req *types.GetProductEndpointReq) (*types.GetProductEndpointResp, error) {
	out, err := l.svcCtx.ProductEndpointRpc.GetProductEndpoint(l.ctx, &productendpoint.GetProductEndpointReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var resp types.GetProductEndpointResp
	_ = copier.CopyWithOption(&resp, out, xcopier.CopierOption)

	return &resp, nil
}
