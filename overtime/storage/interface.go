package storage

import (
	"context"
	"shared/models"
)

type Storage interface {
	Submit(ctx context.Context, overtime models.Overtime) error
}
