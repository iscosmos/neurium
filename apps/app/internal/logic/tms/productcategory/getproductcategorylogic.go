package productcategory

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoryLogic {
	return &GetProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductCategory 查询产品分类
func (l *GetProductCategoryLogic) GetProductCategory(req *types.GetProductCategoryReq) (*types.GetProductCategoryResp, error) {
	out, err := l.svcCtx.ProductCategoryRpc.GetProductCategory(l.ctx, &productcategory.GetProductCategoryReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var resp types.GetProductCategoryResp
	_ = copier.CopyWithOption(&resp, out, xcopier.CopierOption)

	return &resp, nil
}
