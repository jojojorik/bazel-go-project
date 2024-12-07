package api

import (
	"github.com/gorilla/mux"
)

// SetupRouter returns a mux router that handles certain Rest endpoints.
func SetupRouter() *mux.Router {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	router := mux.NewRouter()

	// Register the routes and handlers
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/questions/random", randomQuestionHandler).Methods("GET")
	router.HandleFunc("/questions/{id}/answer", answerHandler).Methods("POST")
	router.HandleFunc("/leaderboard", leaderBoardHandler).Methods("GET")

	return router
}
