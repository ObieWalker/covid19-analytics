package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/models"
)

func main() {

	collection := helper.ConnectDB()
	byteData := readJSONFile("scripts/popDensity/populationData.json")

	models.UpdateCountriesPopulationDensity(collection, byteData)

	return
}

func readJSONFile(filename string) []models.CountryData {
	byteValues, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	}

	var docs []models.CountryData

	err = json.Unmarshal(byteValues, &docs)
	return docs
}
