package product

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/product"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateProduct 创建产品
func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (*types.CreateProductResp, error) {
	var in product.CreateProductReq
	_ = copier.Copy(&in, req)

	out, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.CreateProductResp{Id: out.Id}, nil
}
