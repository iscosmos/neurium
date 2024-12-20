package productschemamodulelogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductSchemaModuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductSchemaModuleLogic {
	return &CreateProductSchemaModuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProductSchemaModule 创建物模型模块
func (l *CreateProductSchemaModuleLogic) CreateProductSchemaModule(in *pb.CreateProductSchemaModuleReq) (*pb.CreateProductSchemaModuleResp, error) {
	var psm model.ProductSchemaModule
	_ = copier.Copy(&psm, in)

	id, err := l.svcCtx.ProductSchemaModuleModel.Create(l.ctx, &psm)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.CreateProductSchemaModuleResp{
		Id: id,
	}, nil
}
