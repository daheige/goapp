package interceptor

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/daheige/gmicro"
	"github.com/daheige/goapp/internal/pkg/ckeys"
	"github.com/daheige/tigago/gutils"
	"github.com/daheige/tigago/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AccessLog access log.
func AccessLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (res interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			// the error format defined by grpc must be used here to return code, desc
			err = status.Errorf(codes.Internal, "%s", "server inner error")
			logger.Info(ctx, "exec panic", map[string]interface{}{
				"reply":       res,
				"trace_error": fmt.Sprintf("%v", r),
				"full_stack":  string(debug.Stack()),
			})
		}
	}()

	t := time.Now()
	clientIP, _ := gmicro.GetGRPCClientIP(ctx)

	// log.Printf("client_ip: %s\n", clientIP)
	// log.Printf("request: %v\n", req)

	// x-request-id
	var requestId string
	if logID := ctx.Value(logger.XRequestID.String()); logID == nil {
		requestId = gutils.Uuid()
	} else {
		requestId, _ = logID.(string)
	}

	ctx = context.WithValue(ctx, logger.XRequestID, requestId)
	ctx = context.WithValue(ctx, logger.ReqClientIP, clientIP)
	ctx = context.WithValue(ctx, logger.RequestMethod, info.FullMethod)
	ctx = context.WithValue(ctx, logger.RequestURI, info.FullMethod)
	ctx = context.WithValue(ctx, ckeys.UserAgent, "grpc-client")
	logger.Info(ctx, "exec begin", map[string]interface{}{
		"client_ip": clientIP,
	})

	res, err = handler(ctx, req)
	ttd := time.Since(t).Milliseconds()
	if err != nil {
		logger.Error(ctx, "exec error", map[string]interface{}{
			"trace_error": err.Error(),
			"exec_time":   ttd,
			"reply":       res,
		})

		return
	}

	logger.Info(ctx, "exec end", map[string]interface{}{
		"exec_time": ttd,
	})

	return
}
