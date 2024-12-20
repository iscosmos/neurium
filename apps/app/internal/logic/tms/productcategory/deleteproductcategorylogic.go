package productcategory

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCategoryLogic {
	return &DeleteProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductCategory 删除产品分类
func (l *DeleteProductCategoryLogic) DeleteProductCategory(req *types.DeleteProductCategoryReq) (*types.DeleteProductCategoryResp, error) {
	if _, err := l.svcCtx.ProductCategoryRpc.DeleteProductCategory(l.ctx, &productcategory.DeleteProductCategoryReq{
		Id: req.Id,
	}); err != nil {
		return nil, err
	}

	return &types.DeleteProductCategoryResp{}, nil
}
