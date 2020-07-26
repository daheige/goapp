package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/daheige/gmicro"
	"github.com/daheige/thinkgo/logger"
	"google.golang.org/grpc"

	"github.com/daheige/goapp/config"
	"github.com/daheige/goapp/internal/rpc/interceptor"
	"github.com/daheige/goapp/internal/rpc/service"
	"github.com/daheige/goapp/pb"
)

var (
	configDir string
	logDir    string
)

func init() {
	flag.StringVar(&configDir, "config_dir", "./", "config dir")
	flag.StringVar(&logDir, "log_dir", "./logs", "log dir")
	flag.Parse()

	// 日志文件设置
	logger.SetLogDir(logDir)
	logger.SetLogFile("go-grpc.log")
	logger.MaxSize(500)
	logger.TraceFileLine(true) //开启文件名和行数追踪

	// 由于logger基于thinkgo/logger又包装了一层，所以这里是3
	logger.InitLogger(3)

	// init config.
	err := config.InitConfig(configDir)
	if err != nil {
		log.Fatalf("init config err: %v", err)
	}
}

func main() {
	defer config.CloseAllDatabase()

	log.Println("rpc start...")
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
		gmicro.WithHandlerFromEndpoint(pb.RegisterGreeterServiceHandlerFromEndpoint),
		gmicro.WithLogger(gmicro.LoggerFunc(log.Printf)),
		gmicro.WithRequestAccess(true),
		gmicro.WithPrometheus(true),
		gmicro.WithGRPCServerOption(grpc.ConnectionTimeout(10*time.Second)),
		gmicro.WithUnaryInterceptor(interceptor.AccessLog),
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

	log.Fatalln(s.StartGRPCAndHTTPServer(config.AppServerConf.GRPCPort))
}

func shutdownFunc() {
	log.Println("server will shutdown")
}
