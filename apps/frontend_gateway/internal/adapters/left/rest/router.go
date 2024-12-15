package rest

import (
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left/rest/handlers"
	"github.com/rs/cors"
)

func (s *server) addRoutes() http.Handler {
	s.l.Message("Setting up http routes")
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://" + s.cfg.NetworkConfig.Frontend.Ip + ":" + strconv.Itoa(s.cfg.NetworkConfig.Frontend.Port)}, // Your frontend URL
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	h := handlers.NewHandler(s.cfg, s.ctx, s.l, s.userApi)
	s.mux.HandleFunc("POST /user/register", h.RegisterUser)
	s.mux.HandleFunc("POST /user/login", h.LoginUser)
	s.mux.HandleFunc("POST /youtube-video/download", h.YoutubeVideoDownload)
	return c.Handler(s.mux)
}
