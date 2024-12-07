package service

import (
	"github.com/jojojorik/bazel-go-project/internal/models"
	"github.com/jojojorik/bazel-go-project/internal/repository"
)

// GetLeaderboard returns a list of users with their score, sorted high to low
func GetLeaderboard() (*[]models.UserScoreResponse, error) {
	leaderboard, err := repository.GetLeaderboard()

	if err != nil {
		return nil, err
	}

	return leaderboard, nil
}
