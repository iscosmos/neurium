package productschemalogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductSchemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductSchemaLogic {
	return &ListProductSchemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListProductSchema 产品物模型列表
func (l *ListProductSchemaLogic) ListProductSchema(in *pb.ListProductSchemaReq) (*pb.ListProductSchemaResp, error) {
	total, pss, err := l.svcCtx.ProductSchemaModel.List(l.ctx, in.PageNum, in.PageSize, in.ModuleId, in.Keyword)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	tree := buildTree(pss)

	return &pb.ListProductSchemaResp{
		Total: total,
		Data:  tree,
	}, nil
}

func buildTree(pss []*model.ProductSchema) []*pb.ListProductSchemaResp_ProductSchema {
	var rows []*pb.ListProductSchemaResp_ProductSchema
	for _, ps := range pss {
		var row pb.ListProductSchemaResp_ProductSchema
		_ = copier.Copy(&row, ps)
		row.Children = make([]*pb.ListProductSchemaResp_ProductSchema, 0)
		rows = append(rows, &row)
	}

	var tree []*pb.ListProductSchemaResp_ProductSchema
	for _, row := range rows {
		for _, r := range rows {
			if r.ParentId == row.Id {
				row.Children = append(row.Children, r)
			}
		}
		if row.ParentId == 0 {
			tree = append(tree, row)
		}
	}

	return tree
}
