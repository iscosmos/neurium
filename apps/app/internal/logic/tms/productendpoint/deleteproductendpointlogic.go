package productendpoint

import (
	"context"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductEndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductEndpointLogic {
	return &DeleteProductEndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductEndpoint 删除产品通信地址
func (l *DeleteProductEndpointLogic) DeleteProductEndpoint(req *types.DeleteProductEndpointReq) (*types.DeleteProductEndpointResp, error) {
	if _, err := l.svcCtx.ProductEndpointRpc.DeleteProductEndpoint(l.ctx, &productendpoint.DeleteProductEndpointReq{Id: req.Id}); err != nil {
		return nil, err
	}
	return &types.DeleteProductEndpointResp{}, nil
}
