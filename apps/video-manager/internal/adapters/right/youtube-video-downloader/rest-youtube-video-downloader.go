package youtubevideodownloader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	"github.com/kkdai/youtube/v2"
)

type RestYoutubeVideoDownloader struct {
	ctx context.Context
	cfg configs.Config
	l   logger.CustomLogger
}

func NewRestVideoDownloader(ctx context.Context, cfg configs.Config, l logger.CustomLogger) YoutubeVideoDownloaderPorts {
	return &RestYoutubeVideoDownloader{ctx, cfg, l}
}

// TODO:(lidor) The library is currently not working Wait for it to fix its self
func (r *RestYoutubeVideoDownloader) DownloadVideoAndWithoutSave(videoURL string, quality string) (*models.VideoInfo, error) {
	r.l.Info("Downloading youtube using Rest.")
	// Initialize the YouTube client
	client := youtube.Client{}

	// Get video information
	video, err := client.GetVideo(videoURL)
	if err != nil {
		return nil, fmt.Errorf("error getting video info: %v", err)
	}

	// Log all available formats to see what qualities are available
	fmt.Println("Available formats:")
	for _, format := range video.Formats {
		fmt.Printf("Quality: %s, Format: %s\n", format.Quality, format.MimeType)
	}

	// Find the correct stream based on the quality
	var stream *youtube.Format
	for _, s := range video.Formats {
		// Try matching by a substring, e.g., `large`, `hd1080`, `medium`
		if strings.Contains(s.Quality, quality) && strings.Contains(s.MimeType, "video/mp4") {
			stream = &s
			break
		}
	}

	// If no matching quality is found, return an error
	if stream == nil {
		return nil, fmt.Errorf("quality '%s' not found. Available qualities are: %v", quality, getAvailableQualities(video.Formats))
	}

	// Create a pipe to stream the video into memory
	pr, pw := io.Pipe()

	// Create a goroutine to stream the video data
	defer pw.Close()

	// Create a stream for the selected video format
	videoStream, _, err := client.GetStream(video, stream)
	if err != nil {
		err = fmt.Errorf("error getting stream for quality %s: %v", quality, err)
		pw.CloseWithError(err)
		return nil, err
	}

	// Copy the stream to the pipe writer
	_, err = io.Copy(pw, videoStream)
	if err != nil {
		err = fmt.Errorf("error downloading video stream: %v", err)
		pw.CloseWithError(err)
		return nil, err
	}

	// Read the pipe into an in-memory buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, pr)
	if err != nil {
		return nil, fmt.Errorf("error copying video stream to buffer: %v", err)
	}

	// Return the buffer with the video data and the video title
	return nil, nil
}

// Helper function to get available qualities as strings
func getAvailableQualities(formats youtube.FormatList) []string {
	qualities := []string{}
	for _, format := range formats {
		qualities = append(qualities, format.Quality)
	}
	return qualities
}
