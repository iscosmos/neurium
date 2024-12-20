package productcategorylogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"gorm.io/gorm"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryLogic {
	return &UpdateProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductCategory 更新产品分类
func (l *UpdateProductCategoryLogic) UpdateProductCategory(in *pb.UpdateProductCategoryReq) (*pb.UpdateProductCategoryResp, error) {
	pc, err := l.svcCtx.ProductCategoryModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.ProductCategoryNotFoundErr
	}

	children, err := l.svcCtx.ProductCategoryModel.GetChildren(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var ids []int64
	for _, child := range children {
		ids = append(ids, int64(child.ID))
	}

	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		pc.Name = in.Name
		if err := model.NewProductCategoryModel(tx).Update(l.ctx, pc); err != nil {
			return err
		}

		if err := model.NewProductCategoryModel(tx).UpdateParentName(l.ctx, ids, pc.ParentName+"/"+pc.Name); err != nil {
			return err
		}

		return nil
	}); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.UpdateProductCategoryResp{}, nil
}
