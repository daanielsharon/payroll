package test

import (
	"context"
	"shared/models"
	"shared/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"attendance/services"
)

func TestAttend_Weekend(t *testing.T) {
	t.Run("should fail on saturday", func(t *testing.T) {
		mockStorage := new(MockStorage)
		service := services.NewService(mockStorage, func() time.Time {
			return time.Date(2025, 6, 14, 0, 0, 0, 0, time.UTC) // saturday
		})

		ctx := context.Background()

		mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(nil)

		err := service.Attend(ctx)
		assert.Error(t, err)
		assert.Equal(t, "unable to submit on weekend", err.Error())
		mockStorage.AssertNotCalled(t, "ClockIn", mock.Anything)
	})

	t.Run("should fail on sunday", func(t *testing.T) {
		mockStorage := new(MockStorage)
		service := services.NewService(mockStorage, func() time.Time {
			return time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC) // sunday
		})

		ctx := context.Background()

		mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(nil)

		err := service.Attend(ctx)
		assert.Error(t, err)
		assert.Equal(t, "unable to submit on weekend", err.Error())
		mockStorage.AssertNotCalled(t, "ClockIn", mock.Anything)
	})
}

func TestAttend_ClockIn(t *testing.T) {
	mockStorage := new(MockStorage)
	mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(nil)
	mockStorage.On("ClockIn", mock.Anything).Return(nil)
	service := services.NewService(mockStorage, utils.GetCurrentTime)

	ctx := context.Background()

	err := service.Attend(ctx)
	assert.NoError(t, err)
	mockStorage.AssertCalled(t, "ClockIn", mock.Anything)
	mockStorage.AssertNotCalled(t, "ClockOut", mock.Anything)
}

func TestAttend_ClockOut(t *testing.T) {
	now := time.Now()
	attendance := &models.Attendance{
		ClockInAt:  &now,
		ClockOutAt: nil,
	}
	mockStorage := new(MockStorage)
	mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(attendance)
	mockStorage.On("ClockOut", mock.Anything, attendance).Return(nil)
	service := services.NewService(mockStorage, utils.GetCurrentTime)
	ctx := context.Background()

	err := service.Attend(ctx)
	assert.NoError(t, err)
	mockStorage.AssertCalled(t, "ClockOut", mock.Anything, attendance)
}

func TestAttend_AlreadyClockedInandOut(t *testing.T) {
	now := time.Now()
	attendance := &models.Attendance{
		ClockInAt:  &now,
		ClockOutAt: &now,
	}
	mockStorage := new(MockStorage)
	mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(attendance)
	service := services.NewService(mockStorage, utils.GetCurrentTime)
	ctx := context.Background()

	err := service.Attend(ctx)
	assert.Error(t, err)
	assert.Equal(t, "user has clocked in and out", err.Error())
	mockStorage.AssertNotCalled(t, "ClockIn", mock.Anything)
	mockStorage.AssertNotCalled(t, "ClockOut", mock.Anything)
}

func TestAttend_NotClockedIn(t *testing.T) {
	attendance := &models.Attendance{
		ClockInAt:  nil,
		ClockOutAt: nil,
	}
	mockStorage := new(MockStorage)
	mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(attendance)
	mockStorage.On("ClockIn", mock.Anything).Return(nil)
	service := services.NewService(mockStorage, utils.GetCurrentTime)
	ctx := context.Background()

	err := service.Attend(ctx)
	assert.Nil(t, err)
	mockStorage.AssertCalled(t, "ClockIn", mock.Anything)
	mockStorage.AssertNotCalled(t, "ClockOut", mock.Anything)
}

func TestCorruptData(t *testing.T) {
	now := time.Now()
	attendance := &models.Attendance{
		ClockInAt:  nil,
		ClockOutAt: &now,
	}
	mockStorage := new(MockStorage)
	mockStorage.On("GetAttendanceByUserId", mock.Anything).Return(attendance)
	mockStorage.On("ClockIn", mock.Anything).Return(nil)
	service := services.NewService(mockStorage, utils.GetCurrentTime)
	ctx := context.Background()

	err := service.Attend(ctx)
	assert.Error(t, err)
	assert.Equal(t, "user has clocked out without clocking in", err.Error())
	mockStorage.AssertNotCalled(t, "ClockIn", mock.Anything)
	mockStorage.AssertNotCalled(t, "ClockOut", mock.Anything)
}
