package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"
)

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to add user.")
	var user models.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	h.uPorts.AddUser(user)
}
