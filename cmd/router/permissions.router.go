package router

import "github.com/Aboagye-Dacosta/shopBackend/cmd/controller"

func (r *Router) initializePermissionsRoutes(c *controller.Controller) {
	r.router.PathPrefix("/permissions")
	r.router.HandleFunc("", c.HttpGetPermissions)
}
