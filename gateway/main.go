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

func reverseProxy(target string) http.HandlerFunc {
	targetURL, _ := url.Parse(strings.TrimSuffix(target, "/"))
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	prefix := targetURL.Path
	if prefix == "" {
		prefix = "/"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		if !strings.HasPrefix(r.URL.Path, "/") {
			r.URL.Path = "/" + r.URL.Path
		}
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	r := router.NewBaseRouter()
	config := config.LoadConfig()
	port := fmt.Sprintf(":%s", config.Server.GatewayPort)
	r.Handle("/payroll/*", reverseProxy(":"+config.Server.PayrollPort))
	r.Handle("/overtime/*", reverseProxy(":"+config.Server.OvertimePort))
	r.Handle("/attendance/*", reverseProxy(":"+config.Server.AttendancePort))
	r.Handle("/reimbursement/*", reverseProxy(":"+config.Server.ReimbursementPort))

	log.Printf("Gateway starting on port %s", port)
	http.ListenAndServe(port, r)
}
