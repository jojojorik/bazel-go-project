package models

// UserScore defines a data entry for a user and score
type UserScore struct {
	User  string
	Score int
}

// UserScoreResponse defines a struct to pass back
type UserScoreResponse struct {
	User       string `bson:"_id" json:"user"`
	TotalScore int    `bson:"total_score" json:"total_score"`
}
