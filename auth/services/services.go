package services

import (
	"context"
	"errors"
	httphelper "shared/http"
	"shared/models"
	"shared/utils"
)

type Service struct {
	client *httphelper.Client
}

func NewService() ServiceInterface {
	return &Service{
		client: httphelper.NewWithServices(),
	}
}

func (s *Service) Login(username string, password string) (map[string]any, error) {
	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}

	var user models.User
	err := s.client.Do(context.Background(), "user", "GET", "/user?username="+username, nil, user)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return nil, errors.New("invalssid password")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"token": token,
	}, nil
}
