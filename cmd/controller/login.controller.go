package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/codes"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	appError "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/utils"
)

// loginUser godoc
// @Summary      Login user
// @Description  Login a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginRequest true "Login data"
// @Success      200  {object} models.AuthResponse
// @Failure      400  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /api/v1/auth/login [post]
func (c *Controller) HttpLoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.LoginRequest

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

	existingUser, err := c.userService.GetByEmail(r.Context(), user.Email)

	if err != nil {
		if appErr, ok := err.(*appError.AppError); ok {
			response := utils.GenErrorResponse(appErr.Entity, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, response); sendErr != nil {
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

	if existingUser == nil {
		resp := utils.GenErrorResponse(entities.USER, http.StatusNotFound, fmt.Errorf("User does not exist"))
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}

		return
	}

	if err := utils.VerifyHash(user.Password, existingUser.Password); err != nil {
		resp := utils.GenAuthResponse(codes.INVALID_PASSWORD, http.StatusBadRequest)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}

		return
	}

	token, err := utils.GenerateJWT(existingUser.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	resp := utils.GenSuccessResponse(entities.USER, codes.LOGIN_SUCCESS, models.AuthResponse{
		Token: token,
		User:  *existingUser,
	})

	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
