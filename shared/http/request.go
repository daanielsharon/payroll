package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRouteParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func DecodeJSON[T any](r *http.Request) (T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}
