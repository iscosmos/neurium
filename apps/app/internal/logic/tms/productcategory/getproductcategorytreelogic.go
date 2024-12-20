package productcategory

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCategoryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoryTreeLogic {
	return &GetProductCategoryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductCategoryTree 查询产品分类树
func (l *GetProductCategoryTreeLogic) GetProductCategoryTree(req *types.GetProductCategoryTreeReq) (*types.GetProductCategoryTreeResp, error) {
	out, err := l.svcCtx.ProductCategoryRpc.GetProductCategoryTree(l.ctx, &productcategory.GetProductCategoryTreeReq{})
	if err != nil {
		return nil, err
	}

	var data []types.ProductCategoryTreeItem
	_ = copier.Copy(&data, out.Data)

	return &types.GetProductCategoryTreeResp{
		Data: data,
	}, nil
}
