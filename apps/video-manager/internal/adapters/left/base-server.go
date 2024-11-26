package left

import "github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/application"

type BaseServer interface {
	ListenAndServe(application.YoutubeDownloaderPorts) error
}
