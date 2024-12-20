package productcategory

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryLogic {
	return &UpdateProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductCategory 更新产品分类
func (l *UpdateProductCategoryLogic) UpdateProductCategory(req *types.UpdateProductCategoryReq) (*types.UpdateProductCategoryResp, error) {
	var in productcategory.UpdateProductCategoryReq
	_ = copier.Copy(&in, req)

	if _, err := l.svcCtx.ProductCategoryRpc.UpdateProductCategory(l.ctx, &in); err != nil {
		return nil, err
	}

	return &types.UpdateProductCategoryResp{}, nil
}
