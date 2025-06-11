package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
