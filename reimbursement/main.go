package main

import (
	"fmt"
	"log"
	"net/http"
	"reimbursement/app"
	"shared/config"
	"shared/constant"
	"shared/tracing"
)

func main() {
	tracing.MustInit(constant.ServiceReimbursement)
	defer tracing.Shutdown()

	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.ReimbursementPort)

	log.Printf("Reimbursement service starting on port %s", port)
	http.ListenAndServe(port, r)
}
