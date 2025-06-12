package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string     `gorm:"size:50;uniqueIndex;not null" json:"username"`
	PasswordHash string     `gorm:"size:255;not null" json:"password_hash"`
	Role         string     `gorm:"size:20;not null;check:role IN ('employee', 'admin')" json:"role"`
	BaseSalary   float64    `gorm:"type:numeric(14,2);not null;check:base_salary >= 0" json:"base_salary"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    *uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedBy    *uuid.UUID `gorm:"type:uuid" json:"updated_by"`

	Reimbursements []Reimbursement `gorm:"foreignKey:UserID" json:"reimbursements"`
	Overtime       []Overtime      `gorm:"foreignKey:UserID" json:"overtime"`
	Attendance     []Attendance    `gorm:"foreignKey:UserID" json:"attendance"`
	OriginalData   *User           `gorm:"-" json:"-"`
}
