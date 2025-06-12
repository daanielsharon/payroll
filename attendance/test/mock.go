package test

import (
	"context"
	"shared/models"

	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetAttendanceByUserId(ctx context.Context) *models.Attendance {
	args := m.Called(ctx)
	if att := args.Get(0); att != nil {
		return att.(*models.Attendance)
	}
	return nil
}

func (m *MockStorage) ClockIn(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockStorage) ClockOut(ctx context.Context, att *models.Attendance) error {
	args := m.Called(ctx, att)
	return args.Error(0)
}
