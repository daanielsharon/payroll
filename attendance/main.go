package main

import (
	"attendance/app"
	"fmt"
	"log"
	"net/http"
	"shared/config"
	"shared/constant"
	"shared/tracing"
)

func main() {
	tracing.MustInit(constant.ServiceAttendance)
	defer tracing.Shutdown()

	r := app.New()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.AttendancePort)

	log.Printf("Attendance service starting on port %s", port)
	http.ListenAndServe(port, r)
}
