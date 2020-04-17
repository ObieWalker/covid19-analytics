package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/models"
	"github.com/ObieWalker/covid19-analytics/dao"
	"github.com/joho/godotenv"
)

//GetAllCountriesRecords ...
func GetAllCountriesRecords()([]models.Country) {
	err := godotenv.Load()
	if os.Getenv("APP_ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading .env file ", err)
		}
	}

	collection := helper.GetCollection()
	records := dao.GetCountriesCollection(collection)

	return records
}

// UpdateCountriesData this should update the country records on the database
func UpdateCountriesData() {
	err := godotenv.Load()
	if os.Getenv("APP_ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading .env file ", err)
		}
	}

	resp, err := http.Get(os.Getenv("COUNTRIES_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	collection := helper.GetCollection()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	dao.UpdateCountriesCollection(collection, string(body))
	return
}
