package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	Submit(ctx context.Context, overtime models.Overtime) error
}
