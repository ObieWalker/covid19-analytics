package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ObieWalker/covid19-analytics/services"
)

// GetAllCountryRecords ...
func GetAllCountryRecords(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("fetching records...")
	records := services.GetAllCountriesRecords()

	json.NewEncoder(w).Encode(records)
	fmt.Println("Records Fetched Completely.")
}
