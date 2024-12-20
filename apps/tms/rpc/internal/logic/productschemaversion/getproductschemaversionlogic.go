package productschemaversionlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSchemaVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSchemaVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSchemaVersionLogic {
	return &GetProductSchemaVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductSchemaVersion 查看物模型版本
func (l *GetProductSchemaVersionLogic) GetProductSchemaVersion(in *pb.GetProductSchemaVersionReq) (*pb.GetProductSchemaVersionResp, error) {
	psv, err := l.svcCtx.ProductSchemaVersionModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var resp pb.GetProductSchemaVersionResp
	_ = copier.Copy(&resp, psv)

	return &resp, nil
}
