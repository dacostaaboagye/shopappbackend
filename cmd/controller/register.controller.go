package controller

import (
	"encoding/json"
	"net/http"

	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/utils"
)

// registerUser godoc
// @Summary      Register user
// @Description  Register a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login body models.RegisterRequest true "Register data"
// @Success      200  {object} models.AuthResponse
// @Failure      400  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /api/v1/auth/register [post]
func (c *Controller) HttpRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response := utils.GenErrorResponse(entities.USER, http.StatusInternalServerError, err)
		if err = utils.SendResponse(w, response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	if err := user.Validate(); err != nil {
		response := &models.Response{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		if err = utils.SendResponse(w, response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		response := &models.Response{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		if err = utils.SendResponse(w, response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	createdUser, err := c.userService.Create(r.Context(), &user)

	if err != nil {
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

	token, err := utils.GenerateJWT(createdUser.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	resp := utils.GenSuccessResponse(entities.USER, http.StatusCreated, models.AuthResponse{
		Token: token,
		User:  *createdUser,
	})

	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
