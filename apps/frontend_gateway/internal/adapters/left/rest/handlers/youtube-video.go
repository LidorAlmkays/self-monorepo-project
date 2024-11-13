package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/dtos/incoming"
)

func (h *Handler) YoutubeVideoDownload(w http.ResponseWriter, r *http.Request) {
	h.l.Info("Received a request to download youtube video")
	// Parse the JSON request body
	var model incoming.DownloadYoutubeVideoDTO
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Log the requested URL for debugging
	h.l.Info("Requested video URL: " + model.Url)

	// Assume you have a function that determines the filename and video path based on the URL
	filename := "test.video.mp4"
	//TODO:(lidor) change the path to the video receive it from the python as well                                                                                                                         // Set this dynamically based on your requirements
	filepath := "C://Users//lidor//Desktop//projects//self-monorepo-project//apps//frontend_gateway//internal//adapters//left//rest//handlers//test.video.mp4" // Adjust the actual path to the video file

	// Open the video file
	file, err := os.Open(filepath)
	if err != nil {
		err = errors.New("Failed to open video file, error: " + err.Error())
		h.l.Error(err)
		http.Error(w, "Failed to open video file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get file information for the modification time
	fileInfo, err := file.Stat()
	if err != nil {
		// Handle error if file.Stat() fails
		err = errors.New("Failed to retrieve file info, error: " + err.Error())
		h.l.Error(err)
		http.Error(w, "Failed to retrieve file info", http.StatusInternalServerError)
		return
	}
	modTime := fileInfo.ModTime()

	w.Header().Set("Access-Control-Allow-Origin", "*") // Replace * with your frontend URL in production
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	// Set headers for file download
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Type", "video/mp4")

	// Stream the file content as the response
	http.ServeContent(w, r, filename, modTime, file)
}
