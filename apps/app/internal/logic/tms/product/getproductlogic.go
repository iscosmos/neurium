package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProduct 产品详情
func (l *GetProductLogic) GetProduct(req *types.GetProductReq) (*types.GetProductResp, error) {
	out, err := l.svcCtx.ProductRpc.GetProduct(l.ctx, &product.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var resp types.GetProductResp
	_ = copier.CopyWithOption(&resp, out, xcopier.CopierOption)

	return &resp, nil
}
