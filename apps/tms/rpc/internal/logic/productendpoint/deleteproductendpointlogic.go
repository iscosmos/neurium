package productendpointlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductEndpointLogic {
	return &DeleteProductEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductEndpoint 删除产品通信地址
func (l *DeleteProductEndpointLogic) DeleteProductEndpoint(in *pb.DeleteProductEndpointReq) (*pb.DeleteProductEndpointResp, error) {
	if err := l.svcCtx.ProductEndpointModel.Delete(l.ctx, in.Id); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.DeleteProductEndpointResp{}, nil
}
