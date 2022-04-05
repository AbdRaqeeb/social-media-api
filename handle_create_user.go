package main

import (
	"encoding/json"
	"net/http"
)

type parameters struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func (config apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	err = userIsEligible(params.Email, params.Password, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := config.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
