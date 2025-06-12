package constant

type contextKey string

const (
	ContextUserID    contextKey = "userID"
	ContextRole      contextKey = "role"
	ContextIPAddress contextKey = "ipAddress"
	ContextRequestID contextKey = "requestID"
	ContextTraceID   contextKey = "traceID"
)
