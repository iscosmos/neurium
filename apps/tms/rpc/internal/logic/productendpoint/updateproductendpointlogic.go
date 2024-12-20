package productendpointlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductEndpointLogic {
	return &UpdateProductEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductEndpoint 更新产品通信地址
func (l *UpdateProductEndpointLogic) UpdateProductEndpoint(in *pb.UpdateProductEndpointReq) (*pb.UpdateProductEndpointResp, error) {
	pe, err := l.svcCtx.ProductEndpointModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	_ = copier.Copy(pe, in)

	if err := l.svcCtx.ProductEndpointModel.Update(l.ctx, pe); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.UpdateProductEndpointResp{}, nil
}
