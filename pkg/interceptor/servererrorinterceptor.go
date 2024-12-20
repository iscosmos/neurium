package interceptor

import (
	"context"
	"errors"
	"github.com/iscosmos/neurium/pkg/code"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)

		if err != nil {
			// 如果是内部自定义错误此处转换成grpc的internal错误码并组合错误信息
			var c code.Code
			if errors.As(err, &c) {
				return resp, status.Newf(codes.Internal, "%d:%s", c.Code, c.Msg).Err()
			}

			// 非自定义错误记录日志
			logx.WithContext(ctx).Error(err)

			return nil, err
		}
		return resp, nil
	}
}
