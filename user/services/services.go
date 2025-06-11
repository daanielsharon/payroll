package services

import (
	"errors"
	"shared/models"
	"user/storage"
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
	if user == nil {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
