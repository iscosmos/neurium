package productcategorylogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductCategoryTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductCategoryTreeLogic {
	return &GetProductCategoryTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductCategoryTree 查询产品分类树
func (l *GetProductCategoryTreeLogic) GetProductCategoryTree(in *pb.GetProductCategoryTreeReq) (*pb.GetProductCategoryTreeResp, error) {
	pcs, err := l.svcCtx.ProductCategoryModel.List(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	tree := buildTree(pcs)

	return &pb.GetProductCategoryTreeResp{
		Data: tree,
	}, nil
}

func buildTree(pcs []*model.ProductCategory) []*pb.GetProductCategoryTreeResp_ProductCategory {
	var rows []*pb.GetProductCategoryTreeResp_ProductCategory
	for _, pc := range pcs {
		var row pb.GetProductCategoryTreeResp_ProductCategory
		_ = copier.Copy(&row, pc)
		row.Children = make([]*pb.GetProductCategoryTreeResp_ProductCategory, 0)
		rows = append(rows, &row)
	}

	var tree []*pb.GetProductCategoryTreeResp_ProductCategory
	for _, row := range rows {
		for _, r := range rows {
			if r.ParentId == row.Id {
				row.Children = append(row.Children, r)
			}
		}
		if row.Id == 1 {
			tree = append(tree, row)
		}
	}

	return tree
}
