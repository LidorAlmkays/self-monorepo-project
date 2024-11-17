package application

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	youtubevideodownloader "github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/right/youtube-video-downloader"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type YoutubeManagerApi struct {
	ctx             context.Context
	cfg             configs.Config
	l               logger.CustomLogger
	videoDownloader youtubevideodownloader.YoutubeVideoDownloaderPorts
}

func NewYoutubeDownloaderApi(ctx context.Context, cfg configs.Config, l logger.CustomLogger, videoDownloader youtubevideodownloader.YoutubeVideoDownloaderPorts) YoutubeDownloaderPorts {
	return &YoutubeManagerApi{
		ctx,
		cfg,
		l,
		videoDownloader,
	}
}

func (y *YoutubeManagerApi) DownloadYoutubeVideo(videoUrl string, quality int) {
	y.videoDownloader.DownloadVideoAndWithoutSave(videoUrl, quality)
}
