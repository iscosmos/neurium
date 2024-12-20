package productschemamodulelogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSchemaModuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSchemaModuleLogic {
	return &GetProductSchemaModuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductSchemaModule 查看物模型模块
func (l *GetProductSchemaModuleLogic) GetProductSchemaModule(in *pb.GetProductSchemaModuleReq) (*pb.GetProductSchemaModuleResp, error) {
	psm, err := l.svcCtx.ProductSchemaModuleModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var resp pb.GetProductSchemaModuleResp
	_ = copier.CopyWithOption(&resp, psm, xcopier.CopierOption)

	return &resp, nil
}
