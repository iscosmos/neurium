package productschemaversionlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSchemaVersionWithLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSchemaVersionWithLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSchemaVersionWithLogic {
	return &GetProductSchemaVersionWithLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductSchemaVersionWith 根据产品和状态查询物模型版本
func (l *GetProductSchemaVersionWithLogic) GetProductSchemaVersionWith(in *pb.GetProductSchemaVersionWithReq) (*pb.GetProductSchemaVersionWithResp, error) {
	psv, err := l.svcCtx.ProductSchemaVersionModel.GetWith(l.ctx, in.ProductId, in.Status)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var resp pb.GetProductSchemaVersionWithResp
	_ = copier.CopyWithOption(&resp, psv, xcopier.CopierOption)

	return &resp, nil
}
