package services

import (
	"context"
	"shared/models"
)

type ServiceInterface interface {
	GetAttendanceByUserId(ctx context.Context) *models.Attendance
	Attend(ctx context.Context) error
}
