package main

import (
	"log"
	"migration/db"

	database "shared/db"
)

func main() {
	gormDB := database.Connect()
	if err := db.Run(gormDB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := db.Seed(gormDB); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}
}
