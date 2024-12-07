package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TriviaQuestion defines a question in a trivia game
type TriviaQuestion struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Question string             `bson:"question" json:"question"`
	Options  []string           `bson:"options,omitempty" json:"options"`
	Answer   string             `bson:"answer" json:"answer"`
}

// PublicTriviaQuestion defines a question that can be publically available
type PublicTriviaQuestion struct {
	ID       int
	Question string
	Options  []string
}

// AnswerRequest defines an answer given by a specific user
type AnswerRequest struct {
	User   string `json:"user"`
	Answer string `json:"answer"`
}

// AnswerResponse defines an answer given by a specific user
type AnswerResponse struct {
	Correct bool
	Message string
	Score   int
}
