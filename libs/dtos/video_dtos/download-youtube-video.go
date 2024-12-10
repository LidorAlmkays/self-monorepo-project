package video_dtos

type DownloadVideoDTO struct {
	VideoUrl     string `json:"videoUrl"`
	VideoQuality int    `json:"videoQuality"`
}
