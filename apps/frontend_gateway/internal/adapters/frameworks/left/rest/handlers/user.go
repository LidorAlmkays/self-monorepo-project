package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
)

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to add user.")

	// Read r.Body into a byte array
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	user := models.UserModel{}
	// Now decode the byte array into the user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.uPorts.AddUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
