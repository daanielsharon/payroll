package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reimbursement/storage"
	"shared/constant"
	"shared/models"
	"time"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}

func (s *Service) Submit(ctx context.Context, reimbursement models.Reimbursement) error {
	return s.storage.Submit(ctx, reimbursement)
}

func (s *Service) GetReimbursementByDate(ctx context.Context, startDate, endDate time.Time) []models.Reimbursement {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceReimbursement))
	ctx, span := tracer.Start(ctx, "GetReimbursementByDate Service")
	defer span.End()

	return s.storage.GetReimbursementByDate(ctx, startDate, endDate)
}

func (s *Service) UpdateReimbursementPayroll(ctx context.Context, startDate, endDate time.Time, payrollRunId uuid.UUID) error {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceReimbursement))
	ctx, span := tracer.Start(ctx, "UpdateReimbursementPayroll Service")
	defer span.End()

	reimbursement := s.GetReimbursementByDate(ctx, startDate, endDate)
	if len(reimbursement) == 0 {
		span.RecordError(errors.New("no reimbursement found"))
		return errors.New("no reimbursement found")
	}

	span.AddEvent("Updating payroll")

	for _, re := range reimbursement {
		data, _ := json.Marshal(re)
		re.OldDataJSON = data

		re.PayrollRunID = &payrollRunId
		err := s.storage.UpdatePayroll(ctx, re)
		if err != nil {
			span.RecordError(fmt.Errorf("error updating payrollRunId: %v", payrollRunId))
		}
	}

	return nil
}
