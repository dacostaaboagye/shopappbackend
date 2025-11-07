package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/internal/database/entities"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	appErrors "github.com/Aboagye-Dacosta/shopBackend/internal/errors"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
	"github.com/gorilla/mux"
)

//get a role

const Role = entities.ROLE

// getRoles godoc
// @Summary      Get role
// @Description  Get a role with it permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Param 		 id   path      string  true  "Role ID"
// @Produce      json
// @Success      200  {object} models.Response{data=models.Role}
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/roles/{id} [get]
func (c *Controller) HttpGetRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	role, err := c.rolesService.GetRole(r.Context(), id)

	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {

			resp := utils.GenErrorResponse(appErr.Entity, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)
		if err := utils.SendResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	resp := utils.GenSuccessResponse(Role, http.StatusOK, role)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

// getRoles godoc
// @Summary      Get roles
// @Description  Get all roles with their permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object} models.Response{data=[]models.Role}
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/roles [get]
func (c *Controller) HttpGetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := c.rolesService.GetAll(r.Context())

	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(Role, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)

		if err := utils.SendResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	resp := utils.GenSuccessResponse(Role, http.StatusOK, roles)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

// createRole godoc
// @Summary      Create role
// @Description  Create a role with its permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      models.RoleRequest  true  "Role id"
// @Success      201  {object} models.Response{data=models.Role}
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/roles [post]
func (c *Controller) HttpCreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.RoleRequest

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}

		return
	}

	createdRole, err := c.rolesService.CreateRole(r.Context(), &role)
	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(Role, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)

		if err := utils.SendResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	resp := utils.GenSuccessResponse(Role, http.StatusOK, createdRole)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

// updateRole godoc
// @Summary      Update role
// @Description  Update a role with its permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      string              true  "Role ID"
// @Param        request  body      models.RoleRequest  true  "Role data"
// @Success      200  {object} models.Response{data=models.Role}
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/roles/{id} [put]
func (c *Controller) HttpUpdateRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var role models.RoleRequest

	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)
		if sendErr := utils.SendResponse(w, resp); sendErr != nil {
			http.Error(w, sendErr.Error(), http.StatusInternalServerError)
		}

		return
	}

	updateRole, err := c.rolesService.UpdateRole(r.Context(), id, &role)

	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(Role, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)

		if err := utils.SendResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	resp := utils.GenSuccessResponse(Role, http.StatusAccepted, updateRole)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

// deleteRole godoc
// @Summary      Delete role
// @Description  Delete a role with its permissions
// @Tags         Roles and Permissions
// @Security     BearerAuth
// @Produce      json
// @Param        id       path      string              true  "Role ID"
// @Success      200  {object} models.Response
// @Failure      400  {object} models.ErrResponse
// @Failure      500  {object} models.ErrResponse
// @Router       /api/v1/roles/{id} [delete]
func (c *Controller) HttpDeleteRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := c.rolesService.DeleteRole(r.Context(), id); err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			resp := utils.GenErrorResponse(Role, appErr.Code, appErr.Err)
			if sendErr := utils.SendResponse(w, resp); sendErr != nil {
				http.Error(w, sendErr.Error(), http.StatusInternalServerError)
			}

			return
		}

		resp := utils.GenErrorResponse(Role, http.StatusInternalServerError, err)

		if err := utils.SendResponse(w, resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		return
	}

	resp := utils.GenSuccessResponse(Role, http.StatusNoContent, id)
	if err := utils.SendResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
