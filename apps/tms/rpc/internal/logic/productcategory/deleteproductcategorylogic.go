package productcategorylogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCategoryLogic {
	return &DeleteProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductCategory 删除产品分类
func (l *DeleteProductCategoryLogic) DeleteProductCategory(in *pb.DeleteProductCategoryReq) (*pb.DeleteProductCategoryResp, error) {
	if _, err := l.svcCtx.ProductCategoryModel.Get(l.ctx, in.Id); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.ProductCategoryNotFoundErr
	}

	// 判断是否存在子分类
	children, err := l.svcCtx.ProductCategoryModel.GetChildren(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}
	if len(children) > 0 {
		return nil, code.ProductCategoryHasChildrenErr
	}

	// 判断是否存在产品
	total, _, err := l.svcCtx.ProductModel.List(l.ctx, &model.ProductListFilter{
		Categories: []int64{in.Id},
	})
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}
	if total > 0 {
		return nil, code.ProductCategoryHasProductErr
	}

	if err := l.svcCtx.ProductCategoryModel.Delete(l.ctx, in.Id); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.DeleteProductCategoryResp{}, nil
}
