package main

import (
	"net/http"
	"strings"
)

func (config apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := config.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
