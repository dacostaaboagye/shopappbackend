package controller

import (
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
)

// getPermissions godoc
// @Summary      Get system permissions
// @Description  Get system permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {object} models.PermissionsResponse
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/permissions [get]
func (c *Controller) HttpGetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := c.permissionService.GetPermissions(r.Context())

	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			resp := utils.GenErrorResponse(appErr.Entity, appErr.Code, appErr.Err)
			if respErr := utils.SendResponse(w, resp); respErr != nil {
				http.Error(w, resp.Message, http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(entities.USER, http.StatusInternalServerError, err)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}
		return
	}

	resp := utils.GenSuccessResponse(entities.PERMISSIONS, http.StatusOK, permissions)

	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
