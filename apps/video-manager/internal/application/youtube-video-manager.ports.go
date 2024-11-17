package application

type YoutubeDownloaderPorts interface {
	DownloadYoutubeVideo(videoUrl string, quality int)
}
