package main

import (
	"net/http"
	"strings"
)


func (config apiConfig) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")

	err := config.dbClient.DeletePost(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}