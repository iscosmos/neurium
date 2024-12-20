package productlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductLogic {
	return &ListProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListProduct 产品列表
func (l *ListProductLogic) ListProduct(in *pb.ListProductReq) (*pb.ListProductResp, error) {
	// 获取该分类Id下的所有子分类id
	categories, err := l.svcCtx.ProductCategoryModel.List(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var nodes []*ProductCategoryNode
	_ = copier.Copy(&nodes, categories)

	descendants := getProductCategory(nodes, in.CategoryId)
	descendants = append(descendants, &ProductCategoryNode{Id: in.CategoryId})

	var ids []int64
	for _, node := range descendants {
		ids = append(ids, node.Id)
	}

	// 构建过滤条件
	var filter model.ProductListFilter
	_ = copier.Copy(&filter, in)
	filter.Categories = ids

	// 列表查询
	total, products, err := l.svcCtx.ProductModel.List(l.ctx, &filter)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var data []*pb.ListProductResp_Product
	_ = copier.CopyWithOption(&data, products, xcopier.CopierOption)

	return &pb.ListProductResp{
		Total: total,
		Data:  data,
	}, nil
}

type ProductCategoryNode struct {
	Id       int64
	ParentId int64
}

func getProductCategoryChildren(nodes []*ProductCategoryNode, parentId int64) []*ProductCategoryNode {
	var children []*ProductCategoryNode
	for _, node := range nodes {
		if node.ParentId == parentId {
			children = append(children, node)
		}
	}
	return children
}

// 递归查询该产品分类Id下是所有子分类
func getProductCategory(nodes []*ProductCategoryNode, parentId int64) []*ProductCategoryNode {
	var descendants []*ProductCategoryNode

	children := getProductCategoryChildren(nodes, parentId)
	for _, child := range children {
		descendants = append(descendants, child)
		descendants = append(descendants, getProductCategory(nodes, child.Id)...)
	}

	return descendants
}
