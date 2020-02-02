package route

import (
	"github.com/deepaksinghkushwah/module-test1/app/controllers"
	"github.com/gorilla/mux"
)

// GetRoutes return mux router
func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	routes = routes.PathPrefix("/api/v1").Subrouter()
	routes.HandleFunc("/hello", controllers.Hello).Methods("GET")
	return routes
}
