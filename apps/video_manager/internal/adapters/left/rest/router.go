package rest

import (
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/left/rest/handlers"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/application"
	"github.com/rs/cors"
)

func (s *server) addRoutes(youtubeDownloader application.YoutubeDownloaderPorts) http.Handler {
	s.l.Message("Setting up http routes")
	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // cores allowed url
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	h := handlers.NewHandler(s.cfg, s.ctx, s.l, youtubeDownloader)
	s.mux.HandleFunc("POST /youtube/download", h.DownloadYoutubeVideo)

	return c.Handler(s.mux)
}
