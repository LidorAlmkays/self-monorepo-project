package rest

import (
	"libs/golang/logger"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/adapters/frameworks/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/ports"
)

type server struct {
	mux     *http.ServeMux
	cfg     configs.Config
	l       logger.CustomLogger
	userApi ports.UserPort
}

func NewServer(cfg configs.Config, l logger.CustomLogger, userApi ports.UserPort) left.BaseServer {
	mux := http.NewServeMux()
	return &server{mux: mux, cfg: cfg, l: l, userApi: userApi}
}

func (s *server) ListenAndServe() error {
	s.addRoutes()
	s.l.Message("Server ready to receive REST requests, on port: " + strconv.Itoa(s.cfg.Server.Port))
	err := http.ListenAndServe(":"+strconv.Itoa(s.cfg.Server.Port), s.mux)
	if err != nil {
		return err
	}
	return nil
}
