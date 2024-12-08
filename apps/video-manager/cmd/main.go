package main

import (
	"context"
	"log"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/left/rest"
	youtubevideodownloader "github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/adapters/right/youtube-video-downloader"
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

// acts like an init function, but doing it this way i can control the program exit code
func setUp() error {

	ctx := context.Background()

	var err error

	//open configs
	cfg, err := configs.SetUpConfig(true)
	if err != nil {
		return err
	}

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.BaseConfig.ProjectName)

	//start http server to talk with frontend
	var s left.BaseServer = rest.NewRestServer(ctx, *cfg, l)

	youtubevideodownloader := youtubevideodownloader.NewRestVideoDownloader(ctx, *cfg, l)
	err = s.ListenAndServe(application.NewYoutubeDownloaderApi(ctx, *cfg, l, youtubevideodownloader))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}
