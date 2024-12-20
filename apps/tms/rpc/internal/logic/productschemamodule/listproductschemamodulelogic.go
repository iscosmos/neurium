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

type ListProductSchemaModuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductSchemaModuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductSchemaModuleLogic {
	return &ListProductSchemaModuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListProductSchemaModule 物模型模块列表
func (l *ListProductSchemaModuleLogic) ListProductSchemaModule(in *pb.ListProductSchemaModuleReq) (*pb.ListProductSchemaModuleResp, error) {
	psms, err := l.svcCtx.ProductSchemaModuleModel.List(l.ctx, in.VersionId)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var data []*pb.ListProductSchemaModuleResp_ProductSchemaModule
	_ = copier.CopyWithOption(&data, psms, xcopier.CopierOption)

	return &pb.ListProductSchemaModuleResp{
		Data: data,
	}, nil
}
