package router

import (
	httphelper "shared/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewBaseRouter returns a new chi router with common middleware applied
func NewBaseRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(httphelper.JSONContentType)
	return r
}
