package rest

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/adapters/frameworks/left/rest/handlers"
)

func (s *server)addRoutes() {
	s.l.Message("Setting up http routes")
	h :=handlers.NewHandler(s.l)
	s.mux.HandleFunc("PUT /user", h.AddUser)
}