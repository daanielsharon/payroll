package shared_context

import (
	"context"
	"shared/constant"
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, constant.ContextUserID, userID)
}

func WithRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, constant.ContextRole, role)
}

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, constant.ContextToken, token)
}

func GetUserID(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(constant.ContextUserID).(string)
	return val, ok
}

func GetRole(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(constant.ContextRole).(string)
	return val, ok
}

func GetToken(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(constant.ContextToken).(string)
	return val, ok
}
