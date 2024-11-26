package incoming

type DownloadYoutubeVideoDTO struct {
	VideoUrl     string `json:"videoUrl"`
	VideoQuality int    `json:"videoQuality"`
}
