package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"size:50;uniqueIndex;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         string    `gorm:"size:20;not null;check:role IN ('user', 'admin')"`
	BaseSalary   float64   `gorm:"type:numeric(14,2);not null;check:base_salary >= 0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	CreatedBy    uuid.UUID `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID `gorm:"type:uuid;not null"`

	Reimbursements []Reimbursement `gorm:"foreignKey:UserID"`
	Overtime       []Overtime      `gorm:"foreignKey:UserID"`
	Attendance     []Attendance    `gorm:"foreignKey:UserID"`
}
