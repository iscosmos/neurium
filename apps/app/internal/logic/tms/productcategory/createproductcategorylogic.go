package productcategory

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productcategory"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductCategoryLogic {
	return &CreateProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateProductCategory 创建产品分类
func (l *CreateProductCategoryLogic) CreateProductCategory(req *types.CreateProductCategoryReq) (*types.CreateProductCategoryResp, error) {
	var in productcategory.CreateProductCategoryReq
	_ = copier.Copy(&in, req)

	in.CreateBy = "admin"

	out, err := l.svcCtx.ProductCategoryRpc.CreateProductCategory(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.CreateProductCategoryResp{Id: out.Id}, nil
}
