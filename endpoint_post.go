package main

import (
	"errors"
	"net/http"
)

func (config apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			config.handleGetPosts(w, r)
		case http.MethodPost:
			config.handleCreatePost(w, r)
		case http.MethodDelete:
			config.handleDeletePost(w, r)
		default:
			respondWithError(w, http.StatusNotFound, errors.New("method not supported"))
	}
}
