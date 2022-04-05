package main

import (
	"errors"
	"net/http"
)

func (config apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			config.handleGetUser(w, r)
		case http.MethodPost:
			config.handleCreateUser(w, r)
		case http.MethodPut:
			config.handleUpdateUser(w, r)
		case http.MethodDelete:
			config.handleDeleteUser(w, r)
		default:
			respondWithError(w, http.StatusNotFound, errors.New("method not supported"))
	}
}