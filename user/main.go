package main

import (
	"fmt"
	"log"
	"net/http"
	"shared/config"
	"shared/constant"
	"shared/tracing"
	"user/app"
)

func main() {
	tracing.MustInit(constant.ServiceUser)
	defer tracing.Shutdown()

	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.UserPort)

	log.Printf("User service starting on port %s", port)
	http.ListenAndServe(port, r)
}
