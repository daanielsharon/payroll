package storage

import (
	"context"
	"shared/models"
)

type Storage interface {
	Submit(ctx context.Context, reimbursement models.Reimbursement) error
}
