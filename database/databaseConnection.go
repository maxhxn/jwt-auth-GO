package database

import (
	//"fmt"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb client initializer
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGO_URL")

	//connecting mongodb client
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(MongoDb))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second)

	defer cancel()

	//checking connection to mongodb by pinging
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Mongodb ping failed %v", err)
	}

	return client
}

// connecting mongodb client
var Client *mongo.Client = DBinstance()

// func for getting collection form database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster-0").Collection(collectionName)
	return collection
}
