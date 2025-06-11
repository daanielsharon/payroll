package main

import (
	"auth/app"
	"fmt"
	"log"
	"net/http"
	"shared/config"
	"shared/constant"
	"shared/tracing"
)

func main() {
	tracing.MustInit(constant.ServiceAuth)
	defer tracing.Shutdown()

	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.AuthPort)

	log.Printf("Auth service starting on port %s", port)
	http.ListenAndServe(port, r)
}
