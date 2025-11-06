package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
)

func RecoverPanic(log *logger.AppLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					ctx := r.Context()
					stack := string(debug.Stack())

					log.ErrLogger.ErrorContext(ctx, "panic recovered",
						"error", err,
						"stack", stack,
					)

					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
