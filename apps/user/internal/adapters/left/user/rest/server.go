package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
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

func (s *server) ListenAndServe(userApi application.UserPort) error {
	handler := s.addRoutes(userApi)
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.BaseConfig.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.BaseConfig.Port), handler)
	if err != nil {
		return err
	}
	return nil
}
