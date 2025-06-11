package services

import "shared/models"

type ServiceInterface interface {
	GetUserByUsername(username string) (*models.User, error)
}
