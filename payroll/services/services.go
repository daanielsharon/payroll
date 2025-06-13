package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"payroll/storage"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"sync"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	storage storage.Storage
	client  *httphelper.Client
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage, client: httphelper.NewWithServices()}
}

func (s *Service) CreatePayrollPeriod(ctx context.Context, payrollPeriod *models.PayrollPeriod) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServicePayroll))
	ctx, span := tracer.Start(ctx, "CreatePayrollPeriod Service")
	defer span.End()

	return s.storage.CreatePayrollPeriod(ctx, payrollPeriod)
}

func (s *Service) GetPayrollPeriodById(ctx context.Context, id uuid.UUID) *models.PayrollPeriod {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServicePayroll))
	ctx, span := tracer.Start(ctx, "GetPayrollPeriodById Service")
	defer span.End()

	return s.storage.GetPayrollPeriodById(ctx, id)
}

func (s *Service) CreatePayrollRun(ctx context.Context, payrollRun models.PayrollRun) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServicePayroll))
	ctx, span := tracer.Start(ctx, "CreatePayrollRun Service")
	defer span.End()

	period := s.GetPayrollPeriodById(ctx, payrollRun.PeriodID)
	if period == nil {
		span.RecordError(errors.New("period not found"))
		return errors.New("period not found")
	}

	err := s.storage.CreatePayrollRun(ctx, &payrollRun)
	if err != nil {
		span.RecordError(err)
		return err
	}

	startDateString := period.StartDate.Format("2006-01-02")
	endDateString := period.EndDate.Format("2006-01-02")

	services := []string{constant.ServiceAttendance, constant.ServiceOvertime, constant.ServiceReimbursement}

	url := "/update-payroll?startDate=" + startDateString + "&endDate=" + endDateString + "&payrollRunId=" + payrollRun.ID.String()

	span.AddEvent("Updating payroll")

	wg := sync.WaitGroup{}
	for _, service := range services {
		wg.Add(1)

		// Create a copy of service for the goroutine to avoid race conditions
		svc := service
		go func() {
			defer wg.Done()

			span := trace.SpanFromContext(ctx)
			_, err := httphelper.DoAndDecode[any](s.client, ctx, svc, http.MethodPost, url, nil)
			if err != nil {
				span.RecordError(fmt.Errorf("error updating payroll for %s, error: %v", svc, err))
				return
			}

			span.AddEvent(fmt.Sprintf("Updated payroll for %s", svc))
		}()
	}

	wg.Wait()

	return nil
}
