package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/frameworks/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/ports"
)

type server struct {
	ctx     context.Context
	mux     *http.ServeMux
	cfg     configs.Config
	l       logger.CustomLogger
	userApi ports.UserPort
}

func NewServer(ctx context.Context, cfg configs.Config, l logger.CustomLogger, userApi ports.UserPort) left.BaseServer {
	mux := http.NewServeMux()
	return &server{mux: mux, cfg: cfg, l: l, userApi: userApi}
}

func (s *server) ListenAndServe() error {
	handler := s.addRoutes()
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.ServiceConfig.Server.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.ServiceConfig.Server.Port), handler)
	if err != nil {
		return err
	}
	return nil
}
