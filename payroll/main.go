package main

import (
	"fmt"
	"log"
	"net/http"
	"payroll/app"
	"shared/config"
)

func main() {
	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.PayrollPort)

	log.Printf("Payroll service starting on port %s", port)
	http.ListenAndServe(port, r)
}
