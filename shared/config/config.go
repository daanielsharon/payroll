package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type Server struct {
	Host              string
	PayrollPort       string
	OvertimePort      string
	AttendancePort    string
	ReimbursementPort string
	GatewayPort       string
}

type ApplicationConfig struct {
	DB     DB
	Server Server
}

func LoadConfig() ApplicationConfig {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("err", err)
		log.Println("Warning: .env file not found")
	}

	host := os.Getenv("HOST")
	return ApplicationConfig{
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		Server: Server{
			Host:              host,
			PayrollPort:       os.Getenv("PAYROLL_PORT"),
			OvertimePort:      os.Getenv("OVERTIME_PORT"),
			AttendancePort:    os.Getenv("ATTENDANCE_PORT"),
			ReimbursementPort: os.Getenv("REIMBURSEMENT_PORT"),
			GatewayPort:       os.Getenv("GATEWAY_PORT"),
		},
	}
}
