package biz

import (
	"github.com/mequq/go-grpc-http-template/config"
	"go.uber.org/zap"
)

// Tenant Repo is a data access interface
type TenantRepoInterface interface {
}

// TenantUsecase is a usecase int
type TenantUsecaseInterface interface {
}

// TenantUsecase is a usecase struct
type TenantUsecase struct {
	logger *zap.Logger
	cfg    *config.ViperConfig
	repo   TenantRepoInterface
}

// NewTenantUsecase creates a new TenantUsecase
func NewTenantUsecase(logger *zap.Logger, cfg *config.ViperConfig, repo TenantRepoInterface) TenantUsecaseInterface {
	return &TenantUsecase{
		logger: logger,
		cfg:    cfg,
		repo:   repo,
	}
}
