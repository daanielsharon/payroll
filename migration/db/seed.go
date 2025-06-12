package db

import (
	"fmt"
	"math/rand"
	"time"

	"shared/constant"
	"shared/models"
	"shared/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func seedUsers(db *gorm.DB, adminID uuid.UUID) error {
	for i := 1; i <= 100; i++ {
		user := models.User{
			ID:           utils.GenerateUUID(),
			Username:     fmt.Sprintf("%s_%s", constant.RoleEmployee, randString(6)),
			PasswordHash: utils.HashPassword("password123"),
			Role:         constant.RoleEmployee,
			BaseSalary:   50000 + float64(i)*100,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			CreatedBy:    &adminID,
			UpdatedBy:    nil,
		}
		if err := db.Create(&user); err.Error != nil {
			return err.Error
		}
	}
	return nil
}

func seedAdmin(db *gorm.DB) (uuid.UUID, error) {
	adminID := utils.GenerateUUID()
	admin := models.User{
		ID:           utils.GenerateUUID(),
		Username:     fmt.Sprintf("%s_%s", constant.RoleAdmin, randString(6)),
		PasswordHash: utils.HashPassword("adminpass"),
		Role:         constant.RoleAdmin,
		BaseSalary:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    &adminID,
		UpdatedBy:    nil,
	}
	if err := db.Create(&admin); err.Error != nil {
		return uuid.Nil, err.Error
	}
	return admin.ID, nil
}

func Seed(db *gorm.DB) error {
	adminID, err := seedAdmin(db)
	if err != nil {
		return err
	}

	return seedUsers(db, adminID)
}
