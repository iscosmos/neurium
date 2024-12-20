package productschemamodule

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemamodule"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductSchemaModuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductSchemaModuleLogic {
	return &CreateProductSchemaModuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateProductSchemaModule 创建物模型模块
func (l *CreateProductSchemaModuleLogic) CreateProductSchemaModule(req *types.CreateProductSchemaModuleReq) (*types.CreateProductSchemaModuleResp, error) {
	var in productschemamodule.CreateProductSchemaModuleReq
	_ = copier.Copy(&in, req)

	in.CreateBy = "system"

	out, err := l.svcCtx.ProductSchemaModuleRpc.CreateProductSchemaModule(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.CreateProductSchemaModuleResp{
		Id: out.Id,
	}, nil
}
