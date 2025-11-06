package controller

import (
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
	"github.com/gorilla/mux"
)

// getUsers godoc
// @Summary      Get users
// @Description  Get registered users
// @Tags         Users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {object} models.UserResponse
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/users [get]
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

// getUserByID godoc
// @Summary      Get user by ID
// @Description  Get a registered user by their ID
// @Tags         Users
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  models.ErrResponse
// @Failure      500  {object}  models.ErrResponse
// @Router       /api/v1/users/{id} [get]
func (c *Controller) HttpGetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, err := c.userService.GetById(r.Context(), id)

	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(appErr.Entity, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}
			return
		}

		resp := utils.GenErrorResponse(entities.USER, http.StatusInternalServerError, err)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}

		return
	}

	resp := utils.GenSuccessResponse(entities.USER, http.StatusOK, user)
	if sendErr := utils.SendResponse(w, resp); sendErr != nil {
		http.Error(w, sendErr.Error(), http.StatusInternalServerError)
	}

}
