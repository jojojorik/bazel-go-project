package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jojojorik/bazel-go-project/internal/models"
	"github.com/jojojorik/bazel-go-project/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello home!")
}

func randomQuestionHandler(w http.ResponseWriter, r *http.Request) {
	question, err := service.GetNewQuestion()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(question)
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	var answerRequest models.AnswerRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&answerRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questionID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	answerResponse, err := service.AnswerTriviaQuestion(answerRequest, questionID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answerResponse)
}

func leaderBoardHandler(w http.ResponseWriter, r *http.Request) {
	leaderboard, err := service.GetLeaderboard()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaderboard)
}
