package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

// ConnectDatabase initializes the MongoDB connection
func ConnectDatabase() {
	// MongoDB URI (you can replace with your MongoDB URI)
	uri := "mongodb://localhost:27017/tudo-thrifting" // Update with your connection URI if needed

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	// Get the database (replace 'mydatabase' with your actual database name)
	Database = Client.Database("mydatabase")

	fmt.Println("MongoDB connection successful!")
}
