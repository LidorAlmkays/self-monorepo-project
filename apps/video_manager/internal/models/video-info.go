package models

// VideoInfo struct to hold the response data
type VideoInfo struct {
	Title    string `json:"title"`
	VideoURL string `json:"video_url"`
	Format   string `json:"format"`
}
