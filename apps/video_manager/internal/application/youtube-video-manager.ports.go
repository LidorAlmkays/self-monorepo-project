package application

import "github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/models"

type YoutubeDownloaderPorts interface {
	DownloadYoutubeVideo(videoUrl string, quality string) (*models.VideoInfo, error)
}
