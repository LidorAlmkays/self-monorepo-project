package rabbitmq

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/ports"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type Server struct {
	ctx     context.Context
	cfg     configs.Config
	l       logger.CustomLogger
	userApi ports.UserPort
}

func NewServer(ctx context.Context, cfg configs.Config, l logger.CustomLogger, userApi ports.UserPort) *Server {
	return &Server{cfg: cfg, l: l, userApi: userApi}
}

func (s *Server) ListenAndServe() error {
	s.l.Message("Server ready to receive User messages requests from rabbitmq, on exchange: " + s.cfg.SharedConfig.Rabbitmq.UserExchangeName)

	return nil
}
