package productendpointlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/jinzhu/copier"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductEndpointLogic {
	return &CreateProductEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProductEndpoint 创建产品通信地址
func (l *CreateProductEndpointLogic) CreateProductEndpoint(in *pb.CreateProductEndpointReq) (*pb.CreateProductEndpointResp, error) {
	var pe model.ProductEndpoint
	_ = copier.Copy(&pe, in)

	id, err := l.svcCtx.ProductEndpointModel.Create(l.ctx, &pe)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.CreateProductEndpointResp{
		Id: id,
	}, nil
}
