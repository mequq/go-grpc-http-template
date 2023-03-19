package app

import (
	"fmt"
	"net"
	"net/http"

	"github.com/mequq/go-grpc-http-template/config"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

// App is the main application object
type App struct {
	logger  *zap.Logger
	cfg     *config.ViperConfig
	host    string
	port    int
	grpcSvr *grpc.Server
	httpSvr *runtime.ServeMux
}

// NewApp creates a new App object
func NewApp(cfg *config.ViperConfig, logger *zap.Logger, grpcSvr *grpc.Server, httpSvr *runtime.ServeMux) *App {
	return &App{
		logger:  logger,
		cfg:     cfg,
		host:    cfg.ServerConfig.Grpc.Host,
		port:    cfg.ServerConfig.Grpc.Port,
		grpcSvr: grpcSvr,
		httpSvr: httpSvr,
	}
}

// Run runs the application
func (a *App) RunGRPC() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.host, a.port))
	if err != nil {
		a.logger.Fatal("failed to start app ", zap.Error(err))
		return err
	}

	if err := a.grpcSvr.Serve(lis); err != nil {
		a.logger.Fatal("failed to start app ", zap.Error(err))
		return err
	}
	return nil
}

// run HTTP server
func (a *App) RunHTTP() error {

	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%d", a.cfg.ServerConfig.Http.Host, a.cfg.ServerConfig.Http.Port),
		//add otl to http requests
		otelhttp.NewHandler(a.httpSvr, "http-server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
			otelhttp.WithPropagators(b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader|b3.B3SingleHeader))),
		),
	); err != nil {
		a.logger.Fatal("failed to start app ", zap.Error(err))
		return err
	}
	return nil

}
