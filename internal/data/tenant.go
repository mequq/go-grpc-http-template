package data

import (
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"go.uber.org/zap"
)

// TenantRepo is a data struct
type TenantRepo struct {
	cfg    *config.ViperConfig
	logger *zap.Logger
	data   *DataSource
}

// NewTenantRepo creates a new TenantRepo
func NewTenantRepo(cfg *config.ViperConfig, logger *zap.Logger, data *DataSource) biz.TenantRepoInterface {
	return &TenantRepo{
		cfg:    cfg,
		logger: logger,
		data:   data,
	}
}
