package services

import (
	"payroll/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}
