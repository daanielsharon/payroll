package services

import (
	"context"
	"reimbursement/storage"
	"shared/models"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) ServiceInterface {
	return &Service{storage: storage}
}

func (s *Service) Submit(ctx context.Context, reimbursement models.Reimbursement) error {
	return s.storage.Submit(ctx, reimbursement)
}
