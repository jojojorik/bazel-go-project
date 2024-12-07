package repository

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var triviaCollection *mongo.Collection
var leaderboardCollection *mongo.Collection

const mongoURI = "mongodb://localhost:27017"

// InitMongoClient connects to a mongodb client
func InitMongoClient() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	triviaCollection = client.Database("TriviaGame").Collection("TriviaQuestions")
	leaderboardCollection = client.Database("TriviaGame").Collection("Leaderboard")
	fmt.Println("Connected to MongoDB!")
}

// GetTriviaCollection returns the MongoDB collection
func GetTriviaCollection() *mongo.Collection {
	return triviaCollection
}

// CloseMongoClient closes the MongoDB connection
func CloseMongoClient() {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB: %v", err)
	}
	fmt.Println("MongoDB connection closed.")
}
