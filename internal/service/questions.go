package service

import (
	"github.com/jojojorik/bazel-go-project/internal/models"
	"github.com/jojojorik/bazel-go-project/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetNewQuestion returns a new question
func GetNewQuestion() (*models.TriviaQuestion, error) {

	triviaQuestion := models.TriviaQuestion{
		Question: "What is the capital of Australia?",
		Options:  []string{"Sydney", "Melbourne", "Canberra", "Perth"},
		Answer:   "Canberra",
	}

	id, err := repository.PersistTriviaQuestion(&triviaQuestion)
	if err != nil {
		return nil, err
	}

	triviaQuestion.ID = id

	return &triviaQuestion, nil
}

// AnswerTriviaQuestion validates the answer given and generates a response
func AnswerTriviaQuestion(answerRequest models.AnswerRequest, questionID primitive.ObjectID) (*models.AnswerResponse, error) {
	triviaQuestion, err := repository.FindTriviaQuestionByID(questionID)

	if err != nil {
		return nil, err
	}

	if answerRequest.Answer == triviaQuestion.Answer {
		err := repository.PersistScore(models.UserScore{
			User:  answerRequest.User,
			Score: 10,
		})

		if err != nil {
			return nil, err
		}

		return &models.AnswerResponse{
			Correct: true,
			Message: "Correct! The answer was indeed " + answerRequest.Answer,
			Score:   10,
		}, nil
	}

	return &models.AnswerResponse{
		Correct: false,
		Message: "Try again!",
		Score:   0,
	}, nil
}
