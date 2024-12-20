package productschemamodulelogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductSchemaModuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSchemaModuleLogic {
	return &DeleteProductSchemaModuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductSchemaModule 删除物模型模块
func (l *DeleteProductSchemaModuleLogic) DeleteProductSchemaModule(in *pb.DeleteProductSchemaModuleReq) (*pb.DeleteProductSchemaModuleResp, error) {
	if err := l.svcCtx.ProductSchemaModuleModel.Delete(l.ctx, in.Id); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.DeleteProductSchemaModuleResp{}, nil
}
