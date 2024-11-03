package main

import (
	"context"
	"log"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/right/userService"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/application"
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

// acts like an init function, but doing it this way i can control the program exit code
func setUp() error {
	ctx := context.Background()
	var err error
	//open configs
	var cfg configs.Config = configs.Config{}
	cfg.SharedConfig, err = libConfigs.GetConfig[libConfigs.SharedConfigs]("./configs/", "shared-configs.yaml")
	if err != nil {
		return err
	}
	cfg.ServiceConfig, err = libConfigs.GetConfig[configs.ServiceConfig]("./configs/", "frontend-gateway.yaml")
	if err != nil {
		return err
	}

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.SharedConfig.FrontendGateway.ProjectName)

	//create project api with the gui
	var userServiceApi userService.UserServiceApi
	userServiceApi, err = userService.NewRestUserService(ctx, l, cfg)
	if err != nil {
		return err
	}
	userApplication := application.NewUserApi(userServiceApi, l)

	//start http server to talk with frontend
	var s left.BaseServer = rest.NewServer(ctx, cfg, l, userApplication)
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
