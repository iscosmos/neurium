package productschemamodule

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSchemaModuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSchemaModuleLogic {
	return &GetProductSchemaModuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductSchemaModule 查询物模型模块
func (l *GetProductSchemaModuleLogic) GetProductSchemaModule(req *types.GetProductSchemaModuleReq) (*types.GetProductSchemaModuleResp, error) {
	out, err := l.svcCtx.ProductSchemaModuleRpc.GetProductSchemaModule(l.ctx, &productschemamodule.GetProductSchemaModuleReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var resp types.GetProductSchemaModuleResp
	_ = copier.Copy(&resp, out)

	return &resp, nil
}
