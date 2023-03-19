package biz

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mequq/go-grpc-http-template/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"go.uber.org/zap"
)

// Tenant Repo is a data access interface
type TenantRepoInterface interface {
	// get tenant
	GetTenant(ctx context.Context, tenant uuid.UUID) error
}

// TenantUsecase is a usecase int
type TenantUsecaseInterface interface {
	// get tenant id
	GetTenant(ctx context.Context, tenant uuid.UUID) error
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

func (t *TenantUsecase) GetTenant(ctx context.Context, tenent uuid.UUID) error {
	ctx, span := otel.Tracer("healthz").Start(ctx, "readiness-db-check")
	defer span.End()
	time.Sleep(time.Second * 1)
	if err := t.repo.GetTenant(ctx, tenent); err != nil {
		span.RecordError(err)

		t.logger.Error("error!", zap.Error(err))
		span.SetStatus(codes.Error, err.Error())
	}

	return nil

}
