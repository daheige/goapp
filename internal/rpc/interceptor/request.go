package interceptor

import (
	"context"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/daheige/gmicro"
	gRuntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/daheige/goapp/pkg/ckeys"
	"github.com/daheige/goapp/pkg/helper"
	"github.com/daheige/goapp/pkg/logger"
)

var (
	xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP       = http.CanonicalHeaderKey("X-Real-IP")
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
		ctx = context.WithValue(ctx, ckeys.XRequestID, gmicro.RndUUIDMd5())
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

// GatewayAccessLog gateway access log.
func GatewayAccessLog(mux *gRuntime.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 捕捉运行时的异常
		defer func() {
			if err := recover(); err != nil {
				logger.Error(r.Context(), "exec panic error", map[string]interface{}{
					"trace_error": string(debug.Stack()),
				})

				// 当http请求发生了recover或异常就直接终止
				http.Error(w, "server inner error", http.StatusInternalServerError)
				return
			}
		}()

		t := time.Now()

		// wrk测试发现log.Println会枷锁，将内容输出到终端的时候，每次都会sync.Mutex
		// lock,unlock操作
		// log.Println("request before")
		// log.Println("request uri: ", r.RequestURI)

		// 设置一些请求的公共参数到上下文上
		reqId := r.Header.Get("x-request-id")
		if reqId == "" {
			reqId = gmicro.RndUUIDMd5()
		}

		// log.Println("log_id: ", reqId)
		// 将requestId 写入当前上下文中
		r = helper.SetValueToHTTPCtx(r, ckeys.XRequestID, reqId)

		// 通过ClientIpWare之后，这里的r.RemoteAddr就是客户端的ip真实地址
		r = helper.SetValueToHTTPCtx(r, ckeys.ClientIP, httpRealIP(r))
		r = helper.SetValueToHTTPCtx(r, ckeys.RequestMethod, r.Method)
		r = helper.SetValueToHTTPCtx(r, ckeys.RequestURI, r.RequestURI)
		r = helper.SetValueToHTTPCtx(r, ckeys.UserAgent, r.Header.Get("User-Agent"))

		logger.Info(r.Context(), "exec begin", nil)

		mux.ServeHTTP(w, r)

		// log.Println("request end")
		logger.Info(r.Context(), "exec end", map[string]interface{}{
			"exec_time": time.Now().Sub(t).Seconds(),
		})
	})
}

// httpRealIP 获取http访问的客户端真实ip地址
// httpRealIP is a function that sets a http.Request's RemoteAddr to the results
// of parsing either the X-Forwarded-For header or the X-Real-IP header (in that
// order).
//
// This middleware should be inserted fairly early in the middleware stack to
// ensure that subsequent layers (e.g., request loggers) which examine the
// RemoteAddr will see the intended value.
//
// You should only use this middleware if you can trust the headers passed to
// you (in particular, the two headers this middleware uses), for example
// because you have placed a reverse proxy like HAProxy or nginx in front of
// chi. If your reverse proxies are configured to pass along arbitrary header
// values from the client, or if you use this middleware without a reverse
// proxy, malicious clients will be able to make you very sad (or, depending on
// how you're using RemoteAddr, vulnerable to an attack of some sort).
func httpRealIP(r *http.Request) string {
	var ip string
	if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}

		ip = xff[:i]
	} else if xRip := r.Header.Get(xRealIP); xRip != "" {
		ip = xRip
	} else {
		ip = r.RemoteAddr
	}

	return ip
}
