package productcategorylogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductCategoryLogic {
	return &CreateProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProductCategory 创建产品分类
func (l *CreateProductCategoryLogic) CreateProductCategory(in *pb.CreateProductCategoryReq) (*pb.CreateProductCategoryResp, error) {
	parent, err := l.svcCtx.ProductCategoryModel.Get(l.ctx, in.ParentId)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.ProductCategoryParentIdNotFoundErr
	}

	id, err := l.svcCtx.ProductCategoryModel.Create(l.ctx, &model.ProductCategory{
		ParentId:   int64(parent.ID),
		Name:       in.Name,
		ParentName: parent.ParentName + "/" + parent.Name,
		CreateBy:   in.CreateBy,
	})
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.CreateProductCategoryResp{
		Id: id,
	}, nil
}
