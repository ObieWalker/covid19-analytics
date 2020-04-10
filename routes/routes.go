package routes

import (
	"github.com/ObieWalker/covid19-analytics/controllers"
	"github.com/gorilla/mux"
)

// UseRoutes ...
func UseRoutes(router *mux.Router) {
	router.HandleFunc("/cov19-records", controllers.GetAllCountryRecords).Methods("GET")
}
