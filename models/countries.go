package models

import (
  "context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"

  "go.mongodb.org/mongo-driver/mongo"
)

var countries []interface{}

func CreateCountriesCollection(collection *mongo.Collection, countriesData string) {
	err := json.Unmarshal([]byte(countriesData), &countries)
	if err != nil {
		log.Fatal(err)
	}
	
	result, err := collection.InsertMany(context.TODO(), countries)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Total Count: ", len(result.InsertedIDs))
}