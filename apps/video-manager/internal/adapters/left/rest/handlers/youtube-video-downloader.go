package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/dtos/incoming"
)

func (h *Handler) DownloadYoutubeVideo(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received request to download youtube video")

	// Parse JSON Body
	var videoInfo incoming.YoutubeVideoDownloadDTO
	if err := json.NewDecoder(r.Body).Decode(&videoInfo); err != nil {
		http.Error(w, "Invalid JSON body, "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := videoInfo.ValidateData(); err != nil {
		http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
		return
	}

	h.youtubeDownloader.DownloadYoutubeVideo(videoInfo.VideoUrl, videoInfo.Quality)
}
