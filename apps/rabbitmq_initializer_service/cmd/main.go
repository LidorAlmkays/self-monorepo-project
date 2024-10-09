package main

import (
	"context"
	"log"
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/rabbitmq_initializer_service/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/rabbitmq_initializer_service/rabbitmq"
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
	cfg.ServiceConfig, err = libConfigs.GetConfig[configs.ServiceConfig]("./configs/", "rabbitmq-initializer-service.yaml")
	if err != nil {
		return err
	}
	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.ServiceConfig.ProjectName)

	//init the rabbitmq connection
	rabbitmqManager := rabbitmq.NewRabbitmqManager(l, cfg, ctx)
	err = rabbitmqManager.Connect()
	if err != nil {
		return err
	}
	err = rabbitmqManager.InitializeProjectRabbitMq(cfg.ServiceConfig.ServicesNames)
	if err != nil {
		return err
	}
	defer rabbitmqManager.CloseConnection()
	return nil
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}
