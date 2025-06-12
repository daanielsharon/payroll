package main

import (
	"fmt"
	"log"
	"net/http"
	"reimbursement/app"
	"shared/config"
)

func main() {
	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.ReimbursementPort)

	log.Printf("Reimbursement service starting on port %s", port)
	http.ListenAndServe(port, r)
}
