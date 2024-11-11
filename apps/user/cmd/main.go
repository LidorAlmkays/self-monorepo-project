package main

import (
	"context"
	"log"
	"os"

	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left/user/rabbitmq"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left/user/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db/mongodb"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application/auth"
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
	cfg.ServiceConfig, err = libConfigs.GetConfig[configs.ServiceConfig]("./configs/", "user-service.yaml")
	if err != nil {
		return err
	}

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.SharedConfig.UserService.ProjectName)

	//starting db connection
	dbConnection := mongodb.NewMongoApi(ctx, cfg.ServiceConfig.Db.Url, cfg.ServiceConfig.Db.Name, l)
	if err != nil {
		return err
	}
	err = dbConnection.StartDbConnection()
	if err != nil {
		return err
	}
	defer dbConnection.CloseDbConnection()

	//TODO:(lidor) move the paper string and number into yaml config
	authPort := auth.NewPepperSaltAuthenticator("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		75, 1, l)

	userApplication := application.NewUserApi(dbConnection, authPort, l)

	systemCh := make(chan error)
	//start http server
	go func() {
		var s left.BaseServer = rest.NewRestServer(ctx, cfg, l)
		err = s.ListenAndServe(userApplication)
		if err != nil {
			l.Error(err)
		}
		systemCh <- err
	}()
	go func() {
		s, err := rabbitmq.NewRabbitmqUserConsumer(l, ctx, cfg)
		if err != nil {
			l.Error(err)
			systemCh <- err
		}
		err = s.ListenAndServe(userApplication)
		if err != nil {
			l.Error(err)
		}
		systemCh <- err
	}()
	err = <-systemCh

	return err
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}
