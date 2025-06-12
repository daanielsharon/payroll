package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	Submit(ctx context.Context, reimbursement models.Reimbursement) error
}
