package youtubevideodownloader

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type RestYoutubeVideoDownloader struct {
	ctx context.Context
	cfg configs.Config
	l   logger.CustomLogger
}

func NewRestVideoDownloader(ctx context.Context, cfg configs.Config, l logger.CustomLogger) YoutubeVideoDownloaderPorts {
	return &RestYoutubeVideoDownloader{ctx, cfg, l}
}

// TODO:(lidor) make logic for downloading video
func (r *RestYoutubeVideoDownloader) DownloadVideoAndWithoutSave(videoUrl string, quality int) {
	r.l.Info("Downloading youtube using Rest.")
}
