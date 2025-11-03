package router

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/middleware"
	"github.com/Aboagye-Dacosta/shopBackend/logger"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func InitRouter(c *controller.Controller, log *logger.AppLogger) *mux.Router {
	root := mux.NewRouter()
	r := root.PathPrefix("/api/v1").Subrouter()

	r.Use(middleware.RecoverPanic(log))
	r.Use(middleware.WithContext)
	r.Use(middleware.RequestLogger(log))

	appRouter := Router{r}
	appRouter.initializeUserRoutes(c)
	appRouter.initializeRegisterRoutes(c)
	appRouter.initializeDocsRoute(root)

	return root
}
