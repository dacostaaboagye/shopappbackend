package controller

import (
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/utils"
)

func (c *Controller) HttpGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAll(r.Context())

	if err != nil {
		// Check if it's an AppError (from our error system)
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(appErr.Entity, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Fallback for unexpected errors
		resp := utils.GenErrorResponse(entities.USER, http.StatusInternalServerError, err)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Success response
	resp := utils.GenSuccessResponse(entities.USER, http.StatusOK, users)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
