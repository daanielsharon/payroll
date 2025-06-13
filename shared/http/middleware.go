package httphelper

import (
	"context"
	"errors"
	"net/http"
	"shared/config"
	"shared/constant"
	shared_context "shared/context"
	"shared/utils"
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
		ctx := r.Context()
		span := trace.SpanFromContext(ctx)

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			return []byte(config.LoadConfig().JWTSecret), nil
		})

		if err != nil || !token.Valid {
			span.RecordError(errors.New("invalid token"))
			JSONResponse(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		role := claims["role"].(string)

		if userID == "" || role == "" {
			span.RecordError(errors.New("invalid token, missing user id or role"))
			JSONResponse(w, http.StatusUnauthorized, "Invalid token", nil)
			return
		}

		ctx = shared_context.WithUserID(ctx, userID)
		ctx = shared_context.WithRole(ctx, role)
		ctx = shared_context.WithToken(ctx, tokenStr)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = strings.Split(r.RemoteAddr, ":")[0]
		}

		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = utils.GenerateRandomUUIDString()
		}
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
