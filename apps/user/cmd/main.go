package main

import (
	"context"
	"log"
	"os"

	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/frameworks/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/frameworks/left/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/frameworks/right/db"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
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

	//starting db connection
	dbConnection, err := db.CreateDbByType(ctx, cfg.Db.Type, cfg.Db.Url, cfg.Db.Name)
	if err != nil {
		return err
	}
	err = dbConnection.StartDbConnection()
	if err != nil {
		return err
	}
	defer dbConnection.CloseDbConnection()

	userApplication := application.NewUserApi(dbConnection)

	//start http server
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
