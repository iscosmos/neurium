package productendpointlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/pkg/xcopier"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductEndpointLogic {
	return &GetProductEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProductEndpoint 查看产品通信地址
func (l *GetProductEndpointLogic) GetProductEndpoint(in *pb.GetProductEndpointReq) (*pb.GetProductEndpointResp, error) {
	pe, err := l.svcCtx.ProductEndpointModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	var resp pb.GetProductEndpointResp
	_ = copier.CopyWithOption(&resp, pe, xcopier.CopierOption)

	return &resp, nil
}
