package db

import (
	"math/rand"
	"time"

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
			ID:           uuid.New(),
			Username:     "user_" + randString(6),
			PasswordHash: utils.HashPassword("password123"),
			Role:         "user",
			BaseSalary:   50000 + float64(i)*100,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			CreatedBy:    adminID,
			UpdatedBy:    adminID,
		}
		if err := db.Create(&user); err.Error != nil {
			return err.Error
		}
	}
	return nil
}

func seedAdmin(db *gorm.DB) (uuid.UUID, error) {
	admin := models.User{
		ID:           uuid.New(),
		Username:     "admin_" + randString(6),
		PasswordHash: utils.HashPassword("adminpass"),
		Role:         "admin",
		BaseSalary:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    uuid.Nil,
		UpdatedBy:    uuid.Nil,
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
