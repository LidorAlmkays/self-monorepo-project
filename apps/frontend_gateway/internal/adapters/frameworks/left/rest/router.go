package rest

import (
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks/left/rest/handlers"
	"github.com/rs/cors"
)

func (s *server) addRoutes() http.Handler {
	s.l.Message("Setting up http routes")
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{s.cfg.Frontend.Url}, // Your frontend URL
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	h := handlers.NewHandler(s.cfg, s.ctx, s.l, s.userApi)
	s.mux.HandleFunc("PUT /user", h.AddUser)
	return c.Handler(s.mux)
}
