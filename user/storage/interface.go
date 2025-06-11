package storage

import (
	"context"
	"shared/models"
)

type Storage interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
