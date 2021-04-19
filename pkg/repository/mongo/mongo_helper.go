package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDatabase is for connect to mongo db from main.go when app started
func ConnectToDatabase() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("connecting to MongoDB failed ... %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("checking client failed ... %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	//collection := client.Database("cafekala").Collection("user")

	return client, nil
}
