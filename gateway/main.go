package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"shared/config"
	"shared/constant"
	httphelper "shared/http"
	"shared/router"
	"shared/tracing"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func reverseProxy(target, name string) http.HandlerFunc {
	targetURL, _ := url.Parse(strings.TrimSuffix(target, "/"))
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req) // keep default behavior

		// Inject OTEL context so downstream services can continue the trace
		otel.GetTextMapPropagator().Inject(req.Context(), propagation.HeaderCarrier(req.Header))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		spanCtx := trace.SpanContextFromContext(r.Context())
		if spanCtx.HasTraceID() {
			log.Printf("[gateway → %s] trace_id=%s", name, spanCtx.TraceID().String())
		}

		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/"+name)

		if !strings.HasPrefix(r.URL.Path, "/") {
			r.URL.Path = "/" + r.URL.Path
		}
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	tracing.MustInit(constant.ServiceGateway)
	defer tracing.Shutdown()

	r := router.NewBaseRouter()
	cfg := config.LoadConfig()
	port := fmt.Sprintf(":%s", cfg.Server.GatewayPort)

	for _, server := range httphelper.RegisteredServers(cfg) {
		r.Handle("/"+server.Name+"/*", otelhttp.NewHandler(
			reverseProxy(server.Target, server.Name),
			server.Name+"-proxy",
		))
	}

	log.Printf("Gateway starting on port %s", port)
	http.ListenAndServe(port, r)
}
