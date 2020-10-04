package logger

import (
	"context"
	"runtime"
	"runtime/debug"

	"github.com/daheige/goapp/pkg/ckeys"
	"github.com/daheige/goapp/pkg/helper"

	"github.com/daheige/thinkgo/logger"
)

/**
{
    "level":"info",
    "time_local":"2019-11-24T20:07:28.472+0800",
    "msg":"exec begin",
    "options":null,
    "ip":"127.0.0.1",
    "plat":"web",
    "request_method":"GET",
    "trace_line":40,
    "request_uri":"/v1/info/123",
    "log_id":"7bb48d0b-2ef4-fc62-0692-40e72db551ef",
    "trace_file":"/web/go/go-proj/app/web/middleware/log.go",
    "ua":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36"
}
*/

func writeLog(ctx context.Context, levelName string, message string, options map[string]interface{}) {
	ua := ckeys.GetStringByCtxKey(ctx, ckeys.UserAgent)
	// 函数调用
	_, file, line, _ := runtime.Caller(2)
	logInfo := map[string]interface{}{
		"request_uri":    ckeys.GetStringByCtxKey(ctx, ckeys.RequestURI),
		"log_id":         ckeys.GetStringByCtxKey(ctx, ckeys.XRequestID),
		"options":        options,
		"ip":             ckeys.GetStringByCtxKey(ctx, ckeys.ClientIP),
		"ua":             ua,
		"plat":           helper.GetDeviceByUa(ua), // 当前设备匹配
		"request_method": ckeys.GetStringByCtxKey(ctx, ckeys.RequestMethod),
		"trace_line":     line,
		"trace_file":     file,
	}

	switch levelName {
	case "info":
		logger.Info(message, logInfo)
	case "debug":
		logger.Debug(message, logInfo)
	case "warn":
		logger.Warn(message, logInfo)
	case "error":
		logger.Error(message, logInfo)
	case "emergency":
		logger.DPanic(message, logInfo)
	default:
		logger.Info(message, logInfo)
	}
}

func Info(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "info", message, context)
}

func Debug(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "debug", message, context)
}

func Warn(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "warn", message, context)
}

func Error(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "error", message, context)
}

// 致命错误或panic捕获
func Emergency(ctx context.Context, message string, context map[string]interface{}) {
	writeLog(ctx, "emergency", message, context)
}

// 异常捕获处理
func Recover(c interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if ctx, ok := c.(context.Context); ok {
				Emergency(ctx, "exec panic", map[string]interface{}{
					"error":       err,
					"error_trace": string(debug.Stack()),
				})

				return
			}

			logger.DPanic("exec panic", map[string]interface{}{
				"error":       err,
				"error_trace": string(debug.Stack()),
			})
		}
	}()
}
