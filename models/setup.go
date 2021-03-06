package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB returns a collection
func ConnectDB() *mongo.Collection {
	// ConnectDB is a helper for connecting to a MongoDB cluster

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://sh3yk0:1234@sh3yk0.byjkm.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("todos")

	return collection
}
