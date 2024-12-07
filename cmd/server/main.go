package main

import (
	"log"
	"net/http"

	"github.com/jojojorik/bazel-go-project/internal/api"
	"github.com/jojojorik/bazel-go-project/internal/repository"
)

func main() {

	repository.InitMongoClient()
	defer repository.CloseMongoClient()

	router := api.SetupRouter()

	log.Println("Server started on :8080")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
