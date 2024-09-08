package handlers

import (
	"net/http"
)


func (h Handler)AddUser(w http.ResponseWriter, r *http.Request)  {
	h.l.Info("Received a request to add user.")
	w.WriteHeader(http.StatusMisdirectedRequest)
	w.Write([]byte ("no way to add user yet"))
}