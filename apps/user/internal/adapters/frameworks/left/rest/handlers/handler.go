package handlers

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/ports"
)

type Handler struct {
	l      logger.CustomLogger
	cfg    configs.Config
	uPorts ports.UserPort
	ctx    context.Context
}

func NewHandler(cfg configs.Config, ctx context.Context, l logger.CustomLogger, uPorts ports.UserPort) *Handler {
	return &Handler{cfg: cfg, ctx: ctx, l: l, uPorts: uPorts}
}
