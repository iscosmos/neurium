package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductLogic {
	return &ListProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListProduct 产品列表
func (l *ListProductLogic) ListProduct(req *types.ListProductReq) (resp *types.ListProductResp, err error) {
	var in product.ListProductReq
	_ = copier.Copy(&in, req)

	out, err := l.svcCtx.ProductRpc.ListProduct(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	var data []types.ProductItem
	_ = copier.CopyWithOption(&data, out.Data, xcopier.CopierOption)

	return &types.ListProductResp{
		Total: out.Total,
		Data:  data,
	}, nil
}
