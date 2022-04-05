package main

import (
	"encoding/json"
	"net/http"
)

type createPostParameter struct {
	UserEmail string `json:"userEmail"`
	Text      string `json:"text"`
}

func (config apiConfig) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	post := createPostParameter{}

	err := decoder.Decode(&post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	createdPost, err := config.dbClient.CreatePost(post.UserEmail, post.Text)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdPost)
}
