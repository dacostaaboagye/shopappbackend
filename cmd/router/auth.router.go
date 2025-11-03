package router

import "github.com/Aboagye-Dacosta/shopBackend/cmd/controller"

func (r *Router) initializeRegisterRoutes(c *controller.Controller) {
	registerRouter := r.router.PathPrefix("/auth").Subrouter()

	registerRouter.HandleFunc("/login", c.HttpLoginUser).Methods("POST")

	registerRouter.HandleFunc("/register", c.HttpRegisterUser).Methods("POST")

}
