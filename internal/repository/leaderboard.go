package repository

import (
	"context"
	"time"

	"github.com/jojojorik/bazel-go-project/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindScoreByUser finds the total score by user
func FindScoreByUser(user string) (*models.UserScore, error) {
	filterScoreByUser := bson.D{{Key: "$match", Value: bson.D{{Key: "user", Value: user}}}}

	sumScoreByUser := bson.D{
		{Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$user"},
				{Key: "total_score", Value: bson.D{{Key: "$sum", Value: "$score"}}},
			},
		},
	}

	// Context with timeout for query execution
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := leaderboardCollection.Aggregate(ctx, mongo.Pipeline{filterScoreByUser, sumScoreByUser})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result models.UserScore
	if cursor.Next(ctx) {
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return &result, nil
}

// GetLeaderboard returns a slice of scores and users
func GetLeaderboard() (*[]models.UserScoreResponse, error) {
	sumScoreByUser := bson.D{
		{Key: "$group",
			Value: bson.D{
				{Key: "_id", Value: "$user"},
				{Key: "total_score", Value: bson.D{{Key: "$sum", Value: "$score"}}},
			},
		},
	}

	sortScoreDescending := bson.D{
		{Key: "$sort",
			Value: bson.D{
				{Key: "total_score", Value: -1},
			},
		},
	}

	// Context with timeout for query execution
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := leaderboardCollection.Aggregate(ctx, mongo.Pipeline{sumScoreByUser, sortScoreDescending})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []models.UserScoreResponse
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// PersistScore saves a user with a score
func PersistScore(userScore models.UserScore) error {
	_, err := leaderboardCollection.InsertOne(context.Background(), userScore)

	return err
}
