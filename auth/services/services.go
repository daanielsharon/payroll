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

	result, err := httphelper.DoAndDecode[models.User](s.client, context.Background(), "user", "GET", "/user?username="+username, nil)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckPasswordHash(password, result.Data.PasswordHash); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(result.Data.ID.String(), result.Data.Role)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"token": token,
	}, nil
}
