//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"github.com/mequq/go-grpc-http-template/internal/data"
	"github.com/mequq/go-grpc-http-template/internal/server"
	"github.com/mequq/go-grpc-http-template/internal/service"
	"github.com/mequq/go-grpc-http-template/pkg/app"
	"go.uber.org/zap"
)

func wireApp(cfg *config.ViperConfig, logger *zap.Logger) (*app.App, error) {
	panic(wire.Build(
		app.NewApp,
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
	))
}
