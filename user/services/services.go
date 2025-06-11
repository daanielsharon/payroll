package services

import (
	"auth/storage"
	"errors"
	"shared/models"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}

func (s *Service) GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}

	user, err := s.storage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
