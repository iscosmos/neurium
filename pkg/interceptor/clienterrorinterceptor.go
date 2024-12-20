package interceptor

import (
	"context"
	"github.com/iscosmos/neurium/pkg/code"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			// 如果为grpc的internal错误码则提取自定义错误信息并重新封装
			if status.Code(err) == codes.Internal {
				sts := status.Convert(err)
				s := sts.Message()

				// 根据规则将错误信息提取后再重新封装
				result := strings.Split(s, ":")
				c, _ := strconv.Atoi(result[0])
				return code.NewCode(c, result[1])
			}

			// 非自定义错误记录日志
			logx.WithContext(ctx).Error(err.Error())
			return err
		}

		return nil
	}
}
