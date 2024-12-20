package productlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProduct 删除产品
func (l *DeleteProductLogic) DeleteProduct(in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	if err := l.svcCtx.ProductModel.Delete(l.ctx, in.Id); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.DeleteProductResp{}, nil
}
