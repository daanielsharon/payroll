package httphelper

import (
	"context"
	"net/http"
	"shared/config"
	"shared/constant"
	shared_context "shared/context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/otel/trace"
)

func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			JSONResponse(w, http.StatusUnauthorized, "Missing token", nil)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return []byte(config.LoadConfig().JWTSecret), nil
		})

		if err != nil || !token.Valid {
			JSONResponse(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		role := claims["role"].(string)

		if userID == "" || role == "" {
			JSONResponse(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		ctx := shared_context.WithUserID(r.Context(), userID)
		ctx = shared_context.WithRole(ctx, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ip := r.Header.Get("X-Forwarded-For")
		requestID := r.Header.Get("X-Request-ID")
		spanCtx := trace.SpanContextFromContext(ctx)
		traceID := ""
		if spanCtx.HasTraceID() {
			traceID = spanCtx.TraceID().String()
		}

		ctx = context.WithValue(ctx, constant.ContextIPAddress, ip)
		ctx = context.WithValue(ctx, constant.ContextRequestID, requestID)
		ctx = context.WithValue(ctx, constant.ContextTraceID, traceID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JSONOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Unsupported Media Type. Only application/json is allowed", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := shared_context.GetRole(r.Context())
		if role != constant.RoleAdmin {
			JSONResponse(w, http.StatusUnauthorized, "User is not admin", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsEmployee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := shared_context.GetRole(r.Context())
		if role != constant.RoleEmployee {
			JSONResponse(w, http.StatusUnauthorized, "User is not employee", nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
