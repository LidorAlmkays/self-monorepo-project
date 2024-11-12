package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/go-playground/validator"
)

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to add user.")
	var user incoming.AddUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.l.Error(errors.New("Invalid request body: " + err.Error()))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, "All fields (name, birthday, username, password) are required", http.StatusBadRequest)
		return
	}

	err = h.uPorts.AddUser(user)
	if err != nil {
		h.l.Error(errors.New("Failed to add user: " + err.Error()))
		http.Error(w, "Failed to add user to the db", http.StatusInternalServerError)
		return
	}
}

func (h Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to authenticate a user.")
	var user incoming.AuthenticateUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.l.Error(errors.New("Invalid request body: " + err.Error()))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userTokenResponse, err := h.uPorts.AuthenticateUser(user)
	if err != nil {
		h.l.Error(errors.New("Failed to create user token: " + err.Error()))
		http.Error(w, "Failed to create user token"+err.Error(), http.StatusUnauthorized)
		return
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(userTokenResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
