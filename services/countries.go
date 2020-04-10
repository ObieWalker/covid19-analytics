package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/models"
	"github.com/joho/godotenv"
)

func UpdateCountriesData() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file ", err)
	}
	
	resp, err := http.Get(os.Getenv("COUNTRIES_URL"))
	
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	collection := helper.ConnectDB()
	models.UpdateCountriesCollection(collection, string(body))
	return
}