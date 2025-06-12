package main

import (
	"fmt"
	"log"
	"net/http"
	"overtime/app"
	"shared/config"
	"shared/constant"
	"shared/tracing"
)

func main() {
	tracing.MustInit(constant.ServiceOvertime)
	defer tracing.Shutdown()

	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.OvertimePort)

	log.Printf("Overtime service starting on port %s", port)
	http.ListenAndServe(port, r)
}
