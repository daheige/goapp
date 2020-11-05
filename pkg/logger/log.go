package logger

import (
	"context"
	"runtime/debug"

	"github.com/daheige/goapp/pkg/ckeys"
	"github.com/daheige/goapp/pkg/helper"
	"github.com/daheige/tigago/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
var (
	// LogEntry 日志接口对象
	LogEntry       logger.Logger
	logStdout      bool             // 日志是否输出到终端
	logEnableColor bool             // 日志是否染色
	logLevel       = zap.DebugLevel // 日志级别
)

// InitLogger 初始化logger实例
// 对于logger option 下面的可以根据实际情况使用
func InitLogger(dir string, filename string) {
	LogEntry = logger.New(
		logger.WithLogDir(dir),                 // 日志目录
		logger.WithLogFilename(filename),       // 日志文件名，默认zap.log
		logger.WithStdout(logStdout),           // 一般生产环境，debug模式进日志输出到终端，建议不输出到stdout
		logger.WithJsonFormat(true),            // json格式化
		logger.WithAddCaller(true),             // 打印行号
		logger.WithCallerSkip(3),               // 如果基于这个Logger包，再包装一次，这个skip = 2,以此类推
		logger.WithEnableColor(logEnableColor), // 日志是否染色，默认不染色
		logger.WithLogLevel(logLevel),          // 设置日志打印最低级别,如果不设置默认为info级别
		logger.WithMaxAge(3),                   // 最大保存3天
		logger.WithMaxSize(200),                // 每个日志文件最大20MB
		logger.WithCompress(false),             // 日志不压缩
		logger.WithEnableCatchStack(true),      // 当使用Panic方法时候是否记录stack信息
	)
}

// WithStdout 日志是否输出到终端
func WithStdout(b bool) {
	logStdout = b
}

// WithEnableColor 日志是否染色
func WithEnableColor(b bool) {
	logEnableColor = b
}

// WithLogLevel 日志级别
func WithLogLevel(level zapcore.Level) {
	logLevel = level
}

func writeLog(ctx context.Context, levelName zapcore.Level, message string, detail map[string]interface{}) {
	ua := ckeys.GetStringByCtxKey(ctx, ckeys.UserAgent)
	logInfo := []zap.Field{
		zap.String(ckeys.RequestURI.String(), ckeys.GetStringByCtxKey(ctx, ckeys.RequestURI)),
		zap.String(ckeys.XRequestID.String(), ckeys.GetStringByCtxKey(ctx, ckeys.XRequestID)),
		zap.Any(ckeys.Detail.String(), detail),
		zap.String(ckeys.ClientIP.String(), ckeys.GetStringByCtxKey(ctx, ckeys.ClientIP)),
		zap.String(ckeys.UserAgent.String(), ua),
		zap.String(ckeys.Plat.String(), helper.GetDeviceByUa(ua)),
		zap.String(ckeys.RequestMethod.String(), ckeys.GetStringByCtxKey(ctx, ckeys.RequestMethod)),
	}

	switch levelName {
	case zapcore.InfoLevel:
		LogEntry.Info(ctx, message, logInfo)
	case zapcore.DebugLevel:
		LogEntry.Debug(ctx, message, logInfo)
	case zapcore.WarnLevel:
		LogEntry.Warn(ctx, message, logInfo)
	case zapcore.ErrorLevel:
		LogEntry.Error(ctx, message, logInfo)
	case zapcore.DPanicLevel:
		LogEntry.DPanic(ctx, message, logInfo)
	default:
		LogEntry.Info(ctx, message, logInfo)
	}
}

// Info info log.
func Info(ctx context.Context, message string, options map[string]interface{}) {
	writeLog(ctx, zapcore.InfoLevel, message, options)
}

// Debug debug log.
func Debug(ctx context.Context, message string, options map[string]interface{}) {
	writeLog(ctx, zapcore.DebugLevel, message, options)
}

// Warn warn log.
func Warn(ctx context.Context, message string, options map[string]interface{}) {
	writeLog(ctx, zapcore.WarnLevel, message, options)
}

// Error error log.
func Error(ctx context.Context, message string, options map[string]interface{}) {
	writeLog(ctx, zapcore.ErrorLevel, message, options)
}

// Emergency 致命错误或panic捕获
func Emergency(ctx context.Context, message string, options map[string]interface{}) {
	writeLog(ctx, zapcore.DPanicLevel, message, options)
}

// Recover 异常捕获处理
func Recover(ctx context.Context, msg ...string) {
	defer func() {
		if err := recover(); err != nil {
			var message = "exec panic"
			if len(msg) > 0 && msg[0] != "" {
				message = msg[0]
			}

			LogEntry.DPanic(ctx, message, []zap.Field{
				zap.String("error_trace", string(debug.Stack())),
				zap.Any("error", err),
			})
		}
	}()
}
