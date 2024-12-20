package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProduct 更新产品
func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (*types.UpdateProductResp, error) {
	var in product.UpdateProductReq
	_ = copier.Copy(&in, req)

	if _, err := l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &in); err != nil {
		return nil, err
	}

	return &types.UpdateProductResp{}, nil
}
