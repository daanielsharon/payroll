package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"shared/config"
	"shared/router"
)

type Server struct {
	Name   string
	Port   string
	Target string
}

func getServers(cfg config.ApplicationConfig) []Server {
	server := cfg.Server
	host := server.Host
	return []Server{
		{
			Name:   "payroll",
			Port:   server.PayrollPort,
			Target: "http://" + host + ":" + server.PayrollPort,
		},
		{
			Name:   "overtime",
			Port:   server.OvertimePort,
			Target: "http://" + host + ":" + server.OvertimePort,
		},
		{
			Name:   "attendance",
			Port:   server.AttendancePort,
			Target: "http://" + host + ":" + server.AttendancePort,
		},
		{
			Name:   "reimbursement",
			Port:   server.ReimbursementPort,
			Target: "http://" + host + ":" + server.ReimbursementPort,
		},
	}
}

func reverseProxy(target, name string) http.HandlerFunc {
	targetURL, _ := url.Parse(strings.TrimSuffix(target, "/"))
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/"+name)
		fmt.Println("url path", r.URL.Path)
		if !strings.HasPrefix(r.URL.Path, "/") {
			r.URL.Path = "/" + r.URL.Path
		}
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	r := router.NewBaseRouter()
	cfg := config.LoadConfig()
	port := fmt.Sprintf(":%s", cfg.Server.GatewayPort)

	for _, server := range getServers(cfg) {
		r.Handle("/"+server.Name+"/*", reverseProxy(server.Target, server.Name))
	}

	log.Printf("Gateway starting on port %s", port)
	http.ListenAndServe(port, r)
}
