package dao

import (
  "context"
	// "fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/ObieWalker/covid19-analytics/helper"
)


func GetCasesAverage(collection *mongo.Collection, id interface{}){
	var doc bson.M
	err := collection.FindOne(context.TODO(), bson.M{ "_id": id}).Decode(&doc);
	if err != nil{
		log.Fatal(err)
	}
	c1 := make(chan float64)
	c2 := make(chan float64)

	go helper.CalculateSliceAverage(c1, doc["weekCases"].(primitive.A), len(doc["weekCases"].(primitive.A)))
	go helper.CalculateSliceAverage(c2, doc["fortnightCases"].(primitive.A), len(doc["fortnightCases"].(primitive.A)))

	weekAverage, fortnightAverage := <- c1, <- c2

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{
		"weekAverage": weekAverage,
		"fortnightAverage": fortnightAverage,
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

