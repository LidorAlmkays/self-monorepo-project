package main

import (
	"libs/golang/logger"
	"log"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/adapters/frameworks/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/adapters/frameworks/left/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/application"
)

// acts like an init function, but doing it this way i can control the program exit code
func setUp() error {
	//open configs
	cfg, err := configs.OpenYamlConfig("./configs/config.yaml")

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger("gateway")
	l.Message("Starting gateway")

	if err != nil {
		return err
	}

	userApplication := application.NewUserApi()

	//start http server
	var s left.BaseServer = rest.NewServer(*cfg, l, userApplication)
	err = s.ListenAndServe()
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
