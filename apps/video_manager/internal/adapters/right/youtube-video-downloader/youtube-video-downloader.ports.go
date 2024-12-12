package youtubevideodownloader

import "github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/models"

type YoutubeVideoDownloaderPorts interface {
	DownloadVideoAndWithoutSave(videoUrl string, quality string) (*models.VideoInfo, error)
	//TODO:(lidor) add to folder DownloadVideoAndSaveToFolder()

}
