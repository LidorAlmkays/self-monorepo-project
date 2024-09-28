package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
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

func (h *Handler) printReceivedMessageBody(r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error(err)
	}
	bodyString := string(bodyBytes)
	h.l.Debug(bodyString)
}
