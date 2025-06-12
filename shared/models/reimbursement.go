package models

import (
	"time"

	"github.com/google/uuid"
)

type Reimbursement struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User         User           `gorm:"foreignKey:UserID"`
	Amount       float64        `gorm:"type:numeric(14,2);not null;check:amount >= 0" json:"amount"`
	Description  string         `gorm:"type:text" json:"description"`
	PayrollRunID *uuid.UUID     `gorm:"type:uuid;index" json:"payroll_run_id"`
	PayrollRun   *PayrollRun    `gorm:"foreignKey:PayrollRunID"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedBy    uuid.UUID      `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedBy    uuid.UUID      `gorm:"type:uuid" json:"updated_by"`
	OriginalData *Reimbursement `gorm:"-" json:"-"`
}
