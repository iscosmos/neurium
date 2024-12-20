package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProduct 删除产品
func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (*types.DeleteProductResp, error) {
	if _, err := l.svcCtx.ProductRpc.DeleteProduct(l.ctx, &product.DeleteProductReq{Id: req.Id}); err != nil {
		return nil, err
	}

	return &types.DeleteProductResp{}, nil
}
