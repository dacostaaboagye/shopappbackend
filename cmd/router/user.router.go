package router

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/middleware"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
)

func (r *Router) initializeUserRoutes(c *controller.Controller) {
	userRouter := r.router.PathPrefix("/users").Subrouter()

	protectRoutes := userRouter.NewRoute().Subrouter()
	protectRoutes.Use(middleware.AuthMiddleWare)
	protectRoutes.HandleFunc("", utils.HandlePermissions("view_users", c.HttpGetUsers)).Methods("GET")

}
