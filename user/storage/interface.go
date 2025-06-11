package storage

import "shared/models"

type Storage interface {
	GetUserByUsername(username string) (*models.User, error)
}
