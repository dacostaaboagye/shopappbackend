package router

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/middleware"
)

func (r *Router) initializeUserRoutes(c *controller.Controller) {
	userRouter := r.router.PathPrefix("/users").Subrouter()

	protectRoutes := userRouter.NewRoute().Subrouter()
	protectRoutes.Use(middleware.AuthMiddleWare)
	protectRoutes.HandleFunc("", c.HttpGetUsers).Methods("GET")

}
