package handlers

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type Handler struct {
	l                 logger.CustomLogger
	cfg               configs.Config
	ctx               context.Context
	youtubeDownloader application.YoutubeDownloaderPorts
}

func NewHandler(cfg configs.Config, ctx context.Context, l logger.CustomLogger, youtubeDownloader application.YoutubeDownloaderPorts) *Handler {
	return &Handler{cfg: cfg, ctx: ctx, l: l, youtubeDownloader: youtubeDownloader}
}
