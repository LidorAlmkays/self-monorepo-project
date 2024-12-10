package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to register user.")

	// Read r.Body into a byte array
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	user := user_dtos.RequestToAddUserDTO{}
	// Now decode the byte array into the user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.uPorts.RegisterUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to login user")
	// Read r.Body into a byte array
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Could not read request body", http.StatusInternalServerError)
		return
	}

	user := user_dtos.AuthenticateUserDTO{}
	// Now decode the byte array into the user struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.uPorts.LoginUser(user)

	if err != nil {
		http.Error(w, "Failed to get token for user: "+err.Error(), http.StatusUnauthorized)
		return
	}
	cookie := &http.Cookie{Name: "session_id",
		Value:   token,
		Secure:  true,
		Expires: time.Now().UTC().Add(1 * time.Hour),
		Path:    "/",
		Domain:  "localhost",
	}
	http.SetCookie(w, cookie)

	response := user_dtos.UserTokenResponseDTO{Token: token}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}
