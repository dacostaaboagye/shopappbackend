package constants

type ctxKey string

const (
	TRACE_ID_KEY   ctxKey = "trace_id"
	REQUEST_ID_KEY ctxKey = "request_id"
	USER_ID_KEY    ctxKey = "user_id"
	PERMISSIONS_KEY ctxKey = "permissions"
	LOGGER_KEY     ctxKey = "logger_key"
)
