package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/daheige/goapp/config"
	"github.com/daheige/goapp/internal/web/routes"
	"github.com/daheige/goapp/pkg/logger"
	"github.com/daheige/tigago/gpprof"
	"github.com/daheige/tigago/monitor"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "go.uber.org/automaxprocs"
)

var (
	configDir string
	wait      time.Duration // 平滑重启的等待时间1s or 1m
)

func init() {
	flag.StringVar(&configDir, "config_dir", "./", "config dir")
	flag.DurationVar(&wait, "graceful_timeout", 3*time.Second, "the server gracefully reload. eg: 15s or 1m")
	flag.Parse()

	// 初始化配置文件
	err := config.InitConfig(configDir)
	if err != nil {
		log.Fatalf("init config err: %v", err)
	}

	// 日志文件设置
	if config.AppServerConf.LogDir == "" {
		config.AppServerConf.LogDir = "./logs"
	}

	// 初始化logger句柄
	logger.InitLogger(config.AppServerConf.LogDir, "go-web.log")

	// 添加prometheus性能监控指标
	prometheus.MustRegister(monitor.WebRequestTotal)
	prometheus.MustRegister(monitor.WebRequestDuration)

	prometheus.MustRegister(monitor.CpuTemp)
	prometheus.MustRegister(monitor.HdFailures)

	// 性能监控的端口port+1000,只能在内网访问
	httpMux := gpprof.New()

	// 添加prometheus metrics处理器
	httpMux.Handle("/metrics", promhttp.Handler())
	gpprof.Run(httpMux, config.AppServerConf.HttpPort+1000)

	// gin mode设置
	switch config.AppServerConf.AppEnv {
	case "local", "dev":
		gin.SetMode(gin.DebugMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	// 这里推荐使用gin.New方法，默认的Default方法的logger,recovery中间件
	// 有些项目也许用不到，另一方面gin recovery 中间件
	// 对于broken pipe存在一些情形无法覆盖到
	// 具体请参考 go-proj/app/web/middleware/log.go#74
	router := gin.New()

	// 加载路由文件中的路由
	routes.WebRoute(router)

	// 服务server设置
	server := &http.Server{
		Handler:           router,
		Addr:              fmt.Sprintf("0.0.0.0:%d", config.AppServerConf.HttpPort),
		IdleTimeout:       20 * time.Second, // tcp idle time
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	// 在独立携程中运行
	log.Println("server run on: ", config.AppServerConf.HttpPort)
	go func() {
		defer logger.Recover(context.Background())

		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.Error(context.Background(), "server close error", map[string]interface{}{
					"trace_error": err.Error(),
				})

				log.Println(err)
				return
			}

			log.Println("server will exit...")
		}
	}()

	// server平滑重启
	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// recivie signal to exit main goroutine
	// window signal
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	// linux signal,please use this in production.
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2, os.Interrupt, syscall.SIGHUP)

	// Block until we receive our signal.
	sig := <-ch

	log.Println("exit signal: ", sig.String())
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// if your application should wait for other services
	// to finalize based on context cancellation.
	go server.Shutdown(ctx) // 在独立的携程中关闭服务器
	<-ctx.Done()

	log.Println("server shutting down")
}
