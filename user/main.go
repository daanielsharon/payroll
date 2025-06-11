package main

import (
	"auth/app"
	"fmt"
	"log"
	"net/http"
	"shared/config"
)

func main() {
	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.UserPort)

	log.Printf("User service starting on port %s", port)
	http.ListenAndServe(port, r)
}
