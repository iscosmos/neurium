package productschemamodule

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductSchemaModuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductSchemaModuleLogic {
	return &ListProductSchemaModuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ListProductSchemaModule 物模型模块列表
func (l *ListProductSchemaModuleLogic) ListProductSchemaModule(req *types.ListProductSchemaModuleReq) (*types.ListProductSchemaModuleResp, error) {
	out, err := l.svcCtx.ProductSchemaModuleRpc.ListProductSchemaModule(l.ctx, &productschemamodule.ListProductSchemaModuleReq{
		VersionId: req.VersionId,
	})
	if err != nil {
		return nil, err
	}

	var data []types.ProductSchemaModuleItem
	_ = copier.Copy(&data, out.Data)

	return &types.ListProductSchemaModuleResp{Data: data}, nil
}
