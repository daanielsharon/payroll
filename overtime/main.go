package main

import (
	"log"
	"net/http"

	"shared/config"
	"shared/router"
)

func main() {
	r := router.NewBaseRouter()
	config := config.LoadConfig()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Overtime Service"}`))
	})

	port := config.Server.OvertimePort
	log.Printf("Overtime service starting on port %s", port)
	http.ListenAndServe(":"+port, r)
}
