package services

import "context"

type ServiceInterface interface {
	Login(ctx context.Context, username, password string) (map[string]any, error)
}
