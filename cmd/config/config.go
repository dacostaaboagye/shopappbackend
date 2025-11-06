package config

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/router"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/service"
	database "github.com/Aboagye-Dacosta/shopBackend/internal/database/db"
	"github.com/Aboagye-Dacosta/shopBackend/internal/database/models"
	"github.com/Aboagye-Dacosta/shopBackend/internal/logger"
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	log := logger.Init()
	db := database.ConnectDB()
	md := models.NewModel(db)
	sr := service.NewService(md)
	ct := controller.NewController(sr)

	return router.InitRouter(ct, log)
}
