package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"shared/config"
	httphelper "shared/http"
	"shared/router"
)

func reverseProxy(target, name string) http.HandlerFunc {
	targetURL, _ := url.Parse(strings.TrimSuffix(target, "/"))
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/"+name)

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

	for _, server := range httphelper.RegisteredServers(cfg) {
		r.Handle("/"+server.Name+"/*", reverseProxy(server.Target, server.Name))
	}

	log.Printf("Gateway starting on port %s", port)
	http.ListenAndServe(port, r)
}
