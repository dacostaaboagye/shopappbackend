package router

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (p *Router) initializeDocsRoute(root *mux.Router) {

	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	root.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	// Serve Swagger UI
	root.PathPrefix("/swagger/").Handler(httpSwagger.Handler())
}
