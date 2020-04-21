package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	err    error
)

//ConnectDB Connect to our mongoDB database
func ConnectDB() *mongo.Collection {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file ", err)
  }
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("ATLAS_URL"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	
	return collection
}

//GetCollection ..
func GetCollection() *mongo.Collection {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	return collection
}

//ErrorResponse Error handler
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

//GetError ...
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
