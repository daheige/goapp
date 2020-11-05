package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/daheige/gmicro"
	"github.com/daheige/goapp/config"
	"github.com/daheige/goapp/internal/rpc/interceptor"
	"github.com/daheige/goapp/internal/rpc/service"
	"github.com/daheige/goapp/pb"
	"github.com/daheige/goapp/pkg/logger"
	"google.golang.org/grpc"
)

var (
	configDir string
)

func init() {
	flag.StringVar(&configDir, "config_dir", "./", "config dir")
	flag.Parse()

	// init config.
	err := config.InitConfig(configDir)
	if err != nil {
		log.Fatalf("init config err: %v", err)
	}

	// 日志文件设置
	if config.AppServerConf.LogDir == "" {
		config.AppServerConf.LogDir = "./logs"
	}

	// 初始化logger句柄
	logger.InitLogger(config.AppServerConf.LogDir, "go-rpc.log")
}

func main() {
	defer config.CloseAllDatabase()

	log.Println("rpc start...")
	log.Println("server pid: ", os.Getppid())

	// add the /test endpoint
	route := gmicro.Route{
		Method:  "GET",
		Pattern: gmicro.PathPattern("test"),
		Handler: func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.Write([]byte("Hello!"))
		},
	}

	// test Option func
	s := gmicro.NewService(
		gmicro.WithRouteOpt(route),
		gmicro.WithShutdownFunc(shutdownFunc),
		gmicro.WithPreShutdownDelay(2*time.Second),
		gmicro.WithShutdownTimeout(5*time.Second),
		gmicro.WithHandlerFromEndpoint(pb.RegisterGreeterServiceHandlerFromEndpoint),
		gmicro.WithLogger(gmicro.LoggerFunc(log.Printf)),
		// gmicro.WithLogger(gmicro.LoggerFunc(gRPCPrintf)), // 定义grpc logger printf
		// gmicro.WithRequestAccess(true),
		gmicro.WithPrometheus(true),
		gmicro.WithGRPCServerOption(grpc.ConnectionTimeout(10*time.Second)),
		gmicro.WithUnaryInterceptor(interceptor.AccessLog), // 自定义访问日志记录
		gmicro.WithGRPCNetwork("tcp"),
		gmicro.WithHTTPHandler(interceptor.GatewayAccessLog), // gateway请求日志记录
	)

	// register grpc service
	pb.RegisterGreeterServiceServer(s.GRPCServer, &service.GreeterService{})

	newRoute := gmicro.Route{
		Method:  "GET",
		Pattern: gmicro.PathPattern("health"),
		Handler: func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		},
	}

	s.AddRoute(newRoute)

	// log.Fatalln(s.StartGRPCAndHTTPServer(config.AppServerConf.GRPCPort))

	// run grpc and http gateway
	log.Fatalln(s.Start(config.AppServerConf.GRPCHttpGatewayPort, config.AppServerConf.GRPCPort))
}

func shutdownFunc() {
	log.Println("server will shutdown")
	logger.Info(context.Background(), "server will shutdown", nil)
}

// gmicro logger printf打印日志函数
func gRPCPrintf(format string, v ...interface{}) {
	logger.Info(context.Background(), fmt.Sprintf(format, v...), nil)
}
