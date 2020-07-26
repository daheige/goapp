package interceptor

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/daheige/goapp/pkg/ckeys"
	"github.com/daheige/goapp/pkg/logger"

	"github.com/daheige/gmicro"
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
				"reply":      res,
				"full_stack": string(debug.Stack()),
			})
		}
	}()

	t := time.Now()
	clientIP, _ := gmicro.GetGRPCClientIP(ctx)

	// log.Printf("client_ip: %s\n", clientIP)
	// log.Printf("request: %v\n", req)

	// request ctx key
	if logID := ctx.Value(ckeys.XRequestID); logID == nil {
		ctx = context.WithValue(ctx, ckeys.XRequestID, gmicro.RndUUID())
	}

	ctx = context.WithValue(ctx, ckeys.ClientIP, clientIP)
	ctx = context.WithValue(ctx, ckeys.RequestMethod, info.FullMethod)
	ctx = context.WithValue(ctx, ckeys.RequestURI, info.FullMethod)
	ctx = context.WithValue(ctx, ckeys.UserAgent, "grpc")

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
