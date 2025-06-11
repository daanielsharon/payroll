package services

type ServiceInterface interface {
	Login(username string, password string) (map[string]any, error)
}
