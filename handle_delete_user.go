package main

import (
	"net/http"
	"strings"
)

func (config apiConfig) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")

	err := config.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

