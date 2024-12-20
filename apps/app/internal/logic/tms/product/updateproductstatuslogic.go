package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStatusLogic {
	return &UpdateProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductStatus 修改发布状态
func (l *UpdateProductStatusLogic) UpdateProductStatus(req *types.UpdateProductStatusReq) (*types.UpdateProductStatusResp, error) {
	var in product.UpdateProductStatusReq
	_ = copier.Copy(&in, req)

	if _, err := l.svcCtx.ProductRpc.UpdateProductStatus(l.ctx, &in); err != nil {
		return nil, err
	}

	return &types.UpdateProductStatusResp{}, nil
}
