package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	pb "github.com/mequq/go-grpc-http-template/api/tenant/v1alpha1"
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/biz"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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

func (t TenantService) GetTenant(ctx context.Context, req *pb.TenantRequest) (*pb.Tenant, error) {
	ctx, span := otel.Tracer("healthz").Start(ctx, "service.tenant")
	defer span.End()
	time.Sleep(time.Second * 1)
	if err := t.tenantUc.GetTenant(ctx, uuid.New()); err != nil {
		t.logger.Error("error!", zap.Error(err))
	}
	span.SetStatus(codes.Ok, "test")
	span.AddEvent("addd")

	return &pb.Tenant{
		Name:        "a",
		Title:       "b",
		Description: "adasd",
		ApiKey:      "asdasd",
	}, nil

}
