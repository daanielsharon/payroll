package db

import (
	"fmt"
	"log"
	"shared/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	log.Println("Database connection established")
	return DB
}
