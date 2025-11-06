package router

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/middleware"
	"github.com/Aboagye-Dacosta/shopBackend/internal/constants"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
)

func (r *Router) initializePermissionsRoutes(c *controller.Controller) {
	permissionRouter := r.router.PathPrefix("/permissions").Subrouter()
	permissionRouter.Use(middleware.AuthMiddleWare)
	permissionRouter.HandleFunc("", utils.HandlePermissions(constants.ManagePermissions, c.HttpGetPermissions))
}
