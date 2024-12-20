package productschemamodule

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSchemaModuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSchemaModuleLogic {
	return &UpdateProductSchemaModuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSchemaModule 更新物模型模块
func (l *UpdateProductSchemaModuleLogic) UpdateProductSchemaModule(req *types.UpdateProductSchemaModuleReq) (*types.UpdateProductSchemaModuleResp, error) {
	var in productschemamodule.UpdateProductSchemaModuleReq
	_ = copier.Copy(&in, req)

	if _, err := l.svcCtx.ProductSchemaModuleRpc.UpdateProductSchemaModule(l.ctx, &in); err != nil {
		return nil, err
	}

	return &types.UpdateProductSchemaModuleResp{}, nil
}
