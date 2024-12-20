package productlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStatusLogic {
	return &UpdateProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductStatus 修改发布状态
func (l *UpdateProductStatusLogic) UpdateProductStatus(in *pb.UpdateProductStatusReq) (*pb.UpdateProductStatusResp, error) {
	p, err := l.svcCtx.ProductModel.Get(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	if p.Status != in.Status {
		p.Status = in.Status
		if err := l.svcCtx.ProductModel.Update(l.ctx, p); err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, code.UnknownErr
		}
	}

	return &pb.UpdateProductStatusResp{}, nil
}
