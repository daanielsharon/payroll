package services

import (
	"context"
	"fmt"
	"payroll/storage"
	"shared/constant"
	"shared/models"

	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}

func (s *Service) CreatePayrollPeriod(ctx context.Context, payrollPeriod models.PayrollPeriod) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServicePayroll))
	ctx, span := tracer.Start(ctx, "CreatePayrollPeriod Service")
	defer span.End()

	return s.storage.CreatePayrollPeriod(ctx, payrollPeriod)
}

func (s *Service) CreatePayrollRun(ctx context.Context, payrollRun models.PayrollRun) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServicePayroll))
	ctx, span := tracer.Start(ctx, "CreatePayrollRun Service")
	defer span.End()

	return s.storage.CreatePayrollRun(ctx, payrollRun)
}
