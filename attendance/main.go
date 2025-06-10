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
	port := fmt.Sprintf(":%s", config.Server.AttendancePort)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Attendance Service"}`))
	})

	log.Printf("Attendance service starting on port %s", config.Server.AttendancePort)
	http.ListenAndServe(port, r)
}
