package main

import (
	"fmt"
	"log"
	"net/http"

	"shared/config"
	"shared/router"
)

func main() {
	r := router.NewBaseRouter()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.ReimbursementPort)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Reimbursement Service"}`))
	})

	log.Printf("Reimbursement service starting on port %s", port)
	http.ListenAndServe(port, r)
}
