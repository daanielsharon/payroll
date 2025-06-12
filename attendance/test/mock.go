package test

import (
	"context"
	"shared/models"
	"time"

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

func (m *MockStorage) GetAttendanceByUserIdAndDate(ctx context.Context, date time.Time) *models.Attendance {
	args := m.Called(ctx, date)
	if att := args.Get(0); att != nil {
		return att.(*models.Attendance)
	}
	return nil
}

func (m *MockStorage) GetAttendanceByDate(ctx context.Context, startDate, endDate time.Time) []models.Attendance {
	args := m.Called(ctx, startDate, endDate)
	if att := args.Get(0); att != nil {
		return att.([]models.Attendance)
	}
	return nil
}

func (m *MockStorage) UpdatePayroll(ctx context.Context, attendance models.Attendance) error {
	args := m.Called(ctx, attendance)
	return args.Error(0)
}

func (m *MockStorage) ClockIn(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockStorage) ClockOut(ctx context.Context, att *models.Attendance) error {
	args := m.Called(ctx, att)
	return args.Error(0)
}
