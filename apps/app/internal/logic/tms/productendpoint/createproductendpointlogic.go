package productendpoint

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/client/productendpoint"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/app/internal/svc"
	"github.com/iscosmos/neurium/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductEndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductEndpointLogic {
	return &CreateProductEndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateProductEndpoint 创建产品通信地址
func (l *CreateProductEndpointLogic) CreateProductEndpoint(req *types.CreateProductEndpointReq) (*types.CreateProductEndpointResp, error) {
	var in productendpoint.CreateProductEndpointReq
	_ = copier.Copy(&in, req)

	out, err := l.svcCtx.ProductEndpointRpc.CreateProductEndpoint(l.ctx, &in)
	if err != nil {
		return nil, err
	}

	return &types.CreateProductEndpointResp{Id: out.Id}, nil
}
