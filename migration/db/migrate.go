package db

import (
	"shared/audit"
	"shared/models"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Attendance{},
		&models.Reimbursement{},
		&models.Overtime{},
		&models.PayrollPeriod{},
		&models.PayrollRun{},
		&audit.Log{},
	)
}
