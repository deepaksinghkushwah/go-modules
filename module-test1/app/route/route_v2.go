package route

import (
	"github.com/deepaksinghkushwah/module-test1/app/controllers"
	"github.com/gorilla/mux"
)

// GetRoutesV2 return mux router
func GetRoutesV2() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	routes = routes.PathPrefix("/api/v2").Subrouter()
	routes.HandleFunc("/hello", controllers.HelloV2).Methods("GET")
	return routes
}
