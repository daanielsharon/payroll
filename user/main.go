package main

import (
	"fmt"
	"log"
	"net/http"
	"shared/config"
	"user/app"
)

func main() {
	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.UserPort)

	log.Printf("User service starting on port %s", port)
	http.ListenAndServe(port, r)
}
