package youtubevideodownloader

type YoutubeVideoDownloaderPorts interface {
	DownloadVideoAndWithoutSave(videoUrl string, quality int)
	//TODO:(lidor) add to folder DownloadVideoAndSaveToFolder()

}
