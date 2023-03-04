package server

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	tenantpb "github.com/mequq/go-grpc-http-template/api/tenant/v1alpha1"
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer(
	cfg *config.ViperConfig,
	logger *zap.Logger,
	tenantSvc *service.TenantService,

) *grpc.Server {
	// grpc server
	srv := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_zap.StreamServerInterceptor(logger),
				grpc_opentracing.StreamServerInterceptor(),
				otelgrpc.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_opentracing.UnaryServerInterceptor(),
				otelgrpc.UnaryServerInterceptor(),
			),
		),
	)
	// register grpc service
	// if in production, you should use reflection only in debug mode
	if !cfg.ServerConfig.Grpc.Production {
		reflection.Register(srv)
	}
	// tenant service
	tenantpb.RegisterTenantServiceServer(srv, tenantSvc)
	health := health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, health)
	return srv
}
