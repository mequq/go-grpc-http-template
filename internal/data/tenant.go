package data

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

func (t *TenantRepo) GetTenant(ctx context.Context, tenant uuid.UUID) error {
	//otel
	_, span := otel.Tracer("healthz").Start(ctx, "readiness-db-check")
	defer span.End()
	//add uuid to attribute
	span.SetAttributes(attribute.String("tenant", tenant.String()))
	time.Sleep(time.Second * 1)
	return errors.New("test")
}
