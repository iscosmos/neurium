package productschemaversion

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productschemaversion"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSchemaVersionWithLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductSchemaVersionWithLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSchemaVersionWithLogic {
	return &GetProductSchemaVersionWithLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetProductSchemaVersionWith 根据产品和状态查询物模型版本
func (l *GetProductSchemaVersionWithLogic) GetProductSchemaVersionWith(req *types.GetProductSchemaVersionWithReq) (*types.GetProductSchemaVersionWithResp, error) {
	var in productschemaversion.GetProductSchemaVersionWithReq
	_ = copier.Copy(&in, req)

	out, err := l.svcCtx.ProductSchemaVersionRpc.GetProductSchemaVersionWith(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	var resp types.GetProductSchemaVersionWithResp
	_ = copier.Copy(&resp, out)

	return &resp, nil
}
