package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type UpdateParameters struct {
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func (config apiConfig) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	decoder := json.NewDecoder(r.Body)

	params := UpdateParameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = userIsEligible(email, params.Password, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := config.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
