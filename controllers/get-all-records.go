package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ObieWalker/covid19-analytics/services"
	"net/http"
)

// GetAllCountryRecords ...
func GetAllCountryRecords(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("fetching records...")
	records := services.GetAllCountriesRecords()

	fmt.Println("Records: ", json.NewEncoder(w).Encode(records))
}
