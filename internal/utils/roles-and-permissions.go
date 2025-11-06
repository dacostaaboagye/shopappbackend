package utils

import (
	"errors"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
)

func HandlePermissions(permission string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rolesAndPermissions := r.Context().Value(logger.PERMISSIONS_KEY).([]string)
		if !CheckPermission(permission, rolesAndPermissions) {
			resp := GenErrorResponse(entities.AUTHORIZATION, http.StatusForbidden, errors.New("you do not have permission to perform this action."))
			if err := SendResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		next(w, r)
	})

}

func CheckPermission(permission string, rolesAndPermissions []string) bool {
	for _, perm := range rolesAndPermissions {
		if perm == permission || perm == "full_access" {
			return true
		}
	}
	return false
}
