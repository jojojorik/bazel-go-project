package repository

import (
	"context"

	"github.com/jojojorik/bazel-go-project/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PersistTriviaQuestion persists a TriviaQuestion and returns any errors that arise
func PersistTriviaQuestion(triviaQuestion *models.TriviaQuestion) (primitive.ObjectID, error) {
	insertOneResult, err := triviaCollection.InsertOne(context.Background(), triviaQuestion)
	id := insertOneResult.InsertedID.(primitive.ObjectID)
	return id, err
}

// FindTriviaQuestionByID retrieves a trivia question by their ID from the MongoDB database
func FindTriviaQuestionByID(questionID primitive.ObjectID) (*models.TriviaQuestion, error) {
	var question models.TriviaQuestion
	filter := bson.D{{Key: "_id", Value: questionID}}

	err := triviaCollection.FindOne(context.Background(), filter).Decode(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// FindAllTriviaQuestions retrieves all users from the MongoDB database
func FindAllTriviaQuestions() ([]models.TriviaQuestion, error) {
	var questions []models.TriviaQuestion

	cursor, err := triviaCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.TriviaQuestion
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		questions = append(questions, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
