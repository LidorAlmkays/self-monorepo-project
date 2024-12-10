package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
)

type server struct {
	ctx     context.Context
	mux     *http.ServeMux
	cfg     configs.Config
	l       logger.CustomLogger
	userApi application.UserPort
}

func NewServer(ctx context.Context, cfg configs.Config, l logger.CustomLogger, userApi application.UserPort) left.BaseServer {
	mux := http.NewServeMux()
	return &server{mux: mux, cfg: cfg, l: l, userApi: userApi}
}

func (s *server) ListenAndServe() error {
	handler := s.addRoutes()
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.BaseConfig.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.BaseConfig.Port), handler)
	if err != nil {
		return err
	}
	return nil
}
