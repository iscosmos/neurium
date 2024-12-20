package productschemamodulelogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSchemaModuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSchemaModuleLogic {
	return &UpdateProductSchemaModuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSchemaModule 更新物模型模块
func (l *UpdateProductSchemaModuleLogic) UpdateProductSchemaModule(in *pb.UpdateProductSchemaModuleReq) (*pb.UpdateProductSchemaModuleResp, error) {
	psm, err := l.svcCtx.ProductSchemaModuleModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	_ = copier.Copy(psm, in)

	if err := l.svcCtx.ProductSchemaModuleModel.Update(l.ctx, psm); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.UpdateProductSchemaModuleResp{}, nil
}
