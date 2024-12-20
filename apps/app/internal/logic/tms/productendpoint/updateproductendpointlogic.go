package productendpoint

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductEndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductEndpointLogic {
	return &UpdateProductEndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductEndpoint 更新产品通信地址
func (l *UpdateProductEndpointLogic) UpdateProductEndpoint(req *types.UpdateProductEndpointReq) (*types.UpdateProductEndpointResp, error) {
	var in productendpoint.UpdateProductEndpointReq
	_ = copier.Copy(&in, req)

	if _, err := l.svcCtx.ProductEndpointRpc.UpdateProductEndpoint(l.ctx, &in); err != nil {
		return nil, err
	}

	return &types.UpdateProductEndpointResp{}, nil
}
