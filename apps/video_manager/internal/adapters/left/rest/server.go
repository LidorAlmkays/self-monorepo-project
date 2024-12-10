package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
)

type server struct {
	ctx context.Context
	mux *http.ServeMux
	cfg configs.Config
	l   logger.CustomLogger
}

func NewRestServer(ctx context.Context, cfg configs.Config, l logger.CustomLogger) left.BaseServer {
	mux := http.NewServeMux()
	return &server{mux: mux, ctx: ctx, cfg: cfg, l: l}
}

func (s *server) ListenAndServe(youtubeDownloader application.YoutubeDownloaderPorts) error {
	handler := s.addRoutes(youtubeDownloader)
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.BaseConfig.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.BaseConfig.Port), handler)
	if err != nil {
		return err
	}

	return nil
}
