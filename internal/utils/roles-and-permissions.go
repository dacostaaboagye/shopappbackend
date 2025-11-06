package utils

import (
	"errors"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/constants"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
)

func HandlePermissions(permission constants.Permission, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissions := r.Context().Value(constants.PERMISSIONS_KEY).([]string)
		permMap := genPermMap(permissions)
		if !CheckPermission(permission, permMap) {
			resp := GenErrorResponse(entities.PERMISSIONS, http.StatusForbidden, errors.New("you do not have permission to perform this action."))
			if err := SendResponse(w, resp); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		next(w, r)
	})
}

func CheckPermission(permission constants.Permission, permMap map[string]struct{}) bool {

	if permMap == nil {
		return false
	}

	if _, ok := permMap[string(permission)]; ok {
		return true
	}

	_, full := permMap[string(constants.FullAccess)]

	return full
}

func genPermMap(perms []string) map[string]struct{} {
	permMap := make(map[string]struct{})
	for _, perm := range perms {
		permMap[perm] = struct{}{}
	}

	return permMap
}
