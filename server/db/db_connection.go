package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	MongoDB := os.Getenv("MONGODB_URI")
	if MongoDB == "" {
		log.Fatal("MONGODB_URI not found in .env file")
	}
	
	clientOptions := options.Client().ApplyURI(MongoDB)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}
	
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB", err)
	}
	
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("DATABASE_NAME not found in .env file")
	}
	
	collection :=  Client.Database(databaseName).Collection(collectionName)
	return collection
	
}
