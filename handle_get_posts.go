package main

import (
	"net/http"
	"strings"
)

func (config apiConfig) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/posts/")

	posts, err := config.dbClient.GetPosts(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
	}

	respondWithJSON(w, http.StatusOK, posts)
}
