package server

import (
	"context"
	"fmt"

	tenantpb "github.com/mequq/go-grpc-http-template/api/tenant/v1alpha1"
	"github.com/mequq/go-grpc-http-template/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// new http server
func NewHttpServer(cfg *config.ViperConfig, logger *zap.Logger) (*runtime.ServeMux, error) {
	// http server
	mux := runtime.NewServeMux()
	endpoint := fmt.Sprintf("%s:%d", cfg.ServerConfig.Grpc.Host, cfg.ServerConfig.Grpc.Port)
	// register http service
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// duble otl
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	if err := tenantpb.RegisterTenantServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return nil, err
	}
	return mux, nil
}
