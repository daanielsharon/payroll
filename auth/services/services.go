package services

import (
	"context"
	"errors"
	"fmt"
	"shared/constant"
	httphelper "shared/http"
	"shared/models"
	"shared/utils"

	"go.opentelemetry.io/otel"
)

type Service struct {
	client *httphelper.Client
}

func NewService() ServiceInterface {
	return &Service{
		client: httphelper.NewWithServices(),
	}
}

func (s *Service) Login(ctx context.Context, username, password string) (map[string]any, error) {
	tracer := otel.Tracer(fmt.Sprintf("%s/service", constant.ServiceAuth))
	ctx, span := tracer.Start(ctx, "Login Service")
	defer span.End()

	if username == "" || password == "" {
		return nil, errors.New("username or password is empty")
	}

	span.AddEvent("Checking user")
	result, err := httphelper.DoAndDecode[models.User](s.client, ctx, "user", "GET", "/user?username="+username, nil)
	if err != nil {
		return nil, err
	}

	span.AddEvent("User found")
	if err := utils.CheckPasswordHash(password, result.Data.PasswordHash); err != nil {
		return nil, errors.New("invalid password")
	}

	span.AddEvent("Password checked")
	token, err := utils.GenerateJWT(result.Data.ID.String(), result.Data.Role)
	if err != nil {
		return nil, err
	}

	span.AddEvent("Token generated")

	return map[string]any{
		"token": token,
	}, nil
}
