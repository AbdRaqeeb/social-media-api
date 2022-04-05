package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbdRaqeeb/social_media_api/internal/database"
)

const address = "localhost:8080"

type apiConfig struct {
	dbClient database.Client
}

func main() {
	db := database.NewClient("db.json")

	config := apiConfig{
		dbClient: db,
	}

	err := db.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/users", config.endpointUsersHandler)
	mux.HandleFunc("/users/", config.endpointUsersHandler)
	mux.HandleFunc("/posts", config.endpointPostsHandler)
	mux.HandleFunc("/posts/", config.endpointPostsHandler)

	srv := http.Server{
		Handler:      mux,
		Addr:         address,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server started on: ", address)

	srv.ListenAndServe()
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if payload != nil {
		response, err := json.Marshal(payload)

		if err != nil {
			log.Println("error marshalling payload", err)
			w.WriteHeader(http.StatusInternalServerError)
			response, _ := json.Marshal(errorBody{
				Error: "Error marshalling payload",
			})

			w.Write(response)
			return
		}

		w.WriteHeader(code)
		w.Write(response)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("don't call respondWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}
