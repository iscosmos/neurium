package productschema

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschema"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductSchemaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductSchemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductSchemaLogic {
	return &ListProductSchemaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListProductSchema 物模型列表
func (l *ListProductSchemaLogic) ListProductSchema(req *types.ListProductSchemaReq) (*types.ListProductSchemaResp, error) {
	var in productschema.ListProductSchemaReq
	_ = copier.Copy(&in, req)

	out, err := l.svcCtx.ProductSchemaRpc.ListProductSchema(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	var data []types.ProductSchemaItem
	_ = copier.CopyWithOption(&data, out.Data, xcopier.CopierOption)

	return &types.ListProductSchemaResp{
		Total: out.Total,
		Data:  data,
	}, nil
}
