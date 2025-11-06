package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
)

const (
	AuthHeader   = "Authorization"
	RequestIdKey = logger.REQUEST_ID_KEY
	UserIDKey    = logger.USER_ID_KEY
	TraceIDKey   = logger.TRACE_ID_KEY
	PermissionsKey = logger.PERMISSIONS_KEY
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AuthHeader)

		if authHeader == "" {
			resp := utils.GenAuthResponse(codes.NO_TOKEN_PROVIDED, http.StatusUnauthorized)
			if err := utils.SendResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			resp := utils.GenAuthResponse(codes.INVALID_TOKEN, http.StatusUnauthorized)
			if err := utils.SendResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		tokenStr := parts[1]

		userid, permissions, err := utils.VerifyJWT(tokenStr)

		if err != nil {
			if ae, ok := err.(*appErrors.AppError); ok {
				resp := utils.GenAuthResponse(ae.Code, http.StatusUnauthorized)
				if sendErr := utils.SendResponse(w, resp); sendErr != nil {
					http.Error(w, sendErr.Error(), http.StatusInternalServerError)
				}
				return
			}
			// Fallback unexpected error: treat as invalid token
			resp := utils.GenAuthResponse(codes.INVALID_TOKEN, http.StatusUnauthorized)

			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userid)
		ctx = context.WithValue(ctx, PermissionsKey, permissions)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
