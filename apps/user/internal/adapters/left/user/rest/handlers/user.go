package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"
	"github.com/go-playground/validator"
)

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to add user.")
	var user models.UserModel

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
