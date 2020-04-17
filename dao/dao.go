package dao

import (
	"context"
	"fmt"
	"encoding/json"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/ObieWalker/covid19-analytics/helper"
	"github.com/ObieWalker/covid19-analytics/models"
)

func getCasesAverage(collection *mongo.Collection, id interface{}){
	var doc bson.M
	err := collection.FindOne(context.TODO(), bson.M{ "_id": id}).Decode(&doc);
	if err != nil{
		log.Fatal(err)
	}
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go helper.CalculateSliceAverage(c1, c2, doc["fortnightCases"].(primitive.A))

	fortnightAverage, weekAverage := <- c1, <- c2

	casesRate := (fortnightAverage-weekAverage)/fortnightAverage

	go helper.OneWeekProjection(c3, weekAverage, casesRate, 0)

	oneWeekProjection:= <- c3

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{
		"fortnightAverage": fortnightAverage,
		"weekAverage": weekAverage,
		"casesRate" : casesRate,
		"oneWeekProjection" : oneWeekProjection,
	}}

	collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil{
		log.Fatal(err)
	}
}

func UpdatePopulationDensity(collection *mongo.Collection, countriesData []models.CountryData){
	fmt.Println("Starting database seed...")
	j := 0
	k := len(countriesData)
	for i := range countriesData {
		doc := countriesData[i]
		filter := bson.M{"country": bson.M{"$eq": doc.Country}}
		update := bson.M{"$set": bson.M{"populationDensity": doc.PopulationDensity}}
		result, _ := collection.UpdateOne(context.TODO(), filter, update)
		if result.MatchedCount > 0 {
			j++
		}
	}
	fmt.Printf("Total successful entries were %v out of %v", j, k)
}

func UpdateWeekCases(collection *mongo.Collection, country string, fortnightChange []float64){
	filter := bson.M{"country": bson.M{"$eq": country}}
	update := bson.M{"$push": bson.M{
		"fortnightCases": bson.M{"$each": fortnightChange, "$slice": -14 }}}
	collection.UpdateOne(context.TODO(), filter, update)
}

func CreateCountriesCollection(collection *mongo.Collection, countriesData string) {
	err := json.Unmarshal([]byte(countriesData), &models.Ctry)
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.InsertMany(context.TODO(), models.Ctry)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Total Count: ", len(result.InsertedIDs))
}

func UpdateCountriesCollection(collection *mongo.Collection, countriesData string) {
	err := json.Unmarshal([]byte(countriesData), &models.Countries)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting database update...")

	j := 0
	for i := range models.Countries {
		var replacedDocument bson.M
		doc := models.Countries[i]
		todaysCasesSli := []float64{doc["todayCases"].(float64)}
		
		filter := bson.M{"country": bson.M{"$eq": doc["country"]}}
		update := bson.D{
			{"$set", bson.D{
				{"cases", doc["cases"]},
				{"todayCases", doc["todayCases"]},
				{"deaths", doc["deaths"]},
				{"todayDeaths", doc["todayDeaths"]},
				{"recovered", doc["recovered"]},
				{"active", doc["active"]},
				{"critical", doc["critical"]},
				{"casesPerOneMillion", doc["casesPerOneMillion"]},
				{"deathsPerOneMillion", doc["deathsPerOneMillion"]},
				{"tests", doc["tests"]},
				{"testsPerOneMillion", doc["testsPerOneMillion"]},
			}},
		{"$push", bson.M{
			"fortnightCases": bson.M{"$each": todaysCasesSli,
		"$slice": -14 }},
		}}

		fmt.Println("just before update")
		err1 := collection.FindOneAndUpdate(
			context.TODO(),
			filter,
			update,
		).Decode(&replacedDocument)

		if err1 != nil{
			log.Fatal(err)
		}

		if replacedDocument["_id"] != nil {
			j++
			go getCasesAverage(collection, replacedDocument["_id"])
		}
	}

	fmt.Printf("%v countries updated", j)
}

//GetCountriesCollection ...
func GetCountriesCollection(collection *mongo.Collection) interface{} {
	var ctx context.Context
	cur, _ := collection.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	resultList := make([]models.Country, 0)
	var result models.Country
	for cur.Next(ctx) {
		err := cur.Decode(&result)
		resultList = append(resultList, result)
		fmt.Println("errors on line 155")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("errors on line 160")
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return resultList
}

