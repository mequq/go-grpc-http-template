// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"github.com/mequq/go-grpc-http-template/internal/data"
	"github.com/mequq/go-grpc-http-template/internal/server"
	"github.com/mequq/go-grpc-http-template/internal/service"
	"github.com/mequq/go-grpc-http-template/pkg/app"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func wireApp(cfg *config.ViperConfig, logger *zap.Logger) (*app.App, error) {
	dataSource, err := data.NewDataSource(logger, cfg)
	if err != nil {
		return nil, err
	}
	tenantRepoInterface := data.NewTenantRepo(cfg, logger, dataSource)
	tenantUsecaseInterface := biz.NewTenantUsecase(logger, cfg, tenantRepoInterface)
	tenantService := service.NewTenantService(logger, cfg, tenantUsecaseInterface)
	grpcServer := server.NewGrpcServer(cfg, logger, tenantService)
	serveMux, err := server.NewHttpServer(cfg, logger)
	if err != nil {
		return nil, err
	}
	appApp := app.NewApp(cfg, logger, grpcServer, serveMux)
	return appApp, nil
}