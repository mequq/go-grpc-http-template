package service

import (
	pb "github.com/mequq/go-grpc-http-template/api/tenant/v1alpha1"
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"go.uber.org/zap"
)

// TenantService is a service struct
type TenantService struct {
	logger *zap.Logger
	cfg    *config.ViperConfig
	pb.UnimplementedTenantServiceServer
	tenantUc biz.TenantUsecaseInterface
}

// NewTenantService creates a new TenantService
func NewTenantService(logger *zap.Logger, cfg *config.ViperConfig, uc biz.TenantUsecaseInterface) *TenantService {
	return &TenantService{
		logger:   logger,
		cfg:      cfg,
		tenantUc: uc,
	}
}
