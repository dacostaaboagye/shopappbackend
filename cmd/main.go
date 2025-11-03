package main

import (
	"net/http"

	"github.com/Aboagye-Dacosta/shopBackend/cmd/config"
	_ "github.com/Aboagye-Dacosta/shopBackend/docs"
	"github.com/Aboagye-Dacosta/shopBackend/internal/env"
)

// @title           Bag Shop Rest API
// @version         1.0
// @description     This is a REST API For Shop management
//
// @contact.name    Solomon Aboagye
// @contact.url     https://github.com/aboagye-dacosta
// @contact.email   dacostaaboagyesolomon@gmail.com
//
// @securityDefinitions.apikey  BearerAuth
// @in header
// @name Authorization
// @description  Type "Bearer Token" in the format **Bearer {token}** to authenticate the request.
//
// @host      localhost:8080
// @BasePath  /api/v1

type application struct {
	port   int
	router http.Handler
}

func main() {
	env.LoadEnv()
	router := config.Setup()
	app := application{
		port:   env.GetIntEnv("PORT", 3000),
		router: router,
	}

	app.Serve()
}
