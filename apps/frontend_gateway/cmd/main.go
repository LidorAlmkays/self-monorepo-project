package main

import (
	"context"
	"log"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks/left/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks/right/rabbitmq"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

// acts like an init function, but doing it this way i can control the program exit code
func setUp() error {
	ctx := context.Background()

	//open configs
	cfg, err := libConfigs.GetConfig[configs.Config]("./configs/")

	if err != nil {
		return err
	}

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.Server.ProjectName)

	//init the rabbitmq connection
	rabbitmqManager := rabbitmq.NewRabbitmqManager(cfg.Rabbitmq.Url, ctx, l)
	err = rabbitmqManager.StartConnection()
	if err != nil {
		return err
	}

	defer rabbitmqManager.CloseConnection()
	// //create project api with the gui
	var userService ports.UserServicePorts
	userService, err = rabbitmqManager.NewUserService()
	if err != nil {
		return err
	}
	userApplication := application.NewUserApi(userService, l)

	//start http server to talk with frontend
	var s left.BaseServer = rest.NewServer(ctx, *cfg, l, userApplication)
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
