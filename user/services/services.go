package services

import (
	"context"
	"errors"
	"fmt"
	"shared/constant"
	"shared/models"
	"user/storage"

	"go.opentelemetry.io/otel"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}

	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceUser))
	ctx, span := tracer.Start(ctx, "GetUserByUsername Service")
	defer span.End()

	span.AddEvent("Checking user")
	user, err := s.storage.GetUserByUsername(ctx, username)
	if user == nil {
		span.AddEvent("User not found")
		return nil, errors.New("user not found")
	}

	if err != nil {
		span.AddEvent("Error getting user")
		return nil, err
	}

	span.AddEvent("User found")
	return user, nil
}
