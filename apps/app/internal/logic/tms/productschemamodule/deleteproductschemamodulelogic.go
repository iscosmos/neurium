package productschemamodule

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductSchemaModuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSchemaModuleLogic {
	return &DeleteProductSchemaModuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductSchemaModule 删除物模型模块
func (l *DeleteProductSchemaModuleLogic) DeleteProductSchemaModule(req *types.DeleteProductSchemaModuleReq) (*types.DeleteProductSchemaModuleResp, error) {
	_, err := l.svcCtx.ProductSchemaModuleRpc.DeleteProductSchemaModule(l.ctx, &productschemamodule.DeleteProductSchemaModuleReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DeleteProductSchemaModuleResp{}, nil
}
