package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey string

func WithContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := r.Header.Get("X-Trace-ID")

		if traceId == "" {
			traceId = uuid.NewString()
		}

		requestId := uuid.NewString()
		ctx := context.WithValue(r.Context(), TraceIDKey, traceId)
		ctx = context.WithValue(ctx, RequestIdKey, requestId)

		w.Header().Set("X-Trace-ID", traceId)
		w.Header().Set("X-Request-ID", requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
