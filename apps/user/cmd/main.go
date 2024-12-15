package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/LidorAlmkays/self-monorepo-project/libs/enums"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
	"github.com/LidorAlmkays/self-monorepo-project/libs/validators"
	"github.com/go-playground/validator"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/left/user/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db/mongodb"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application/auth"
)

type ProgramFlags struct {
	Mode string `validate:"required,programmode"`
}

var programFlags ProgramFlags

func init() {
	flag.StringVar(&programFlags.Mode, "Mode", "development", "This flags changes the program mode")
	flag.Parse()
	programFlags.Mode = strings.ToLower(programFlags.Mode)
	validate := validator.New()
	validate.RegisterValidation("programmode", validators.ProgramModeValidator)
	err := validate.Struct(programFlags)
	if err != nil {
		panic(err)
	}
}

// acts like an init function, but doing it this way i can control the program exit code
func setUp() error {

	ctx := context.Background()

	var err error
	//open configs
	cfg, err := configs.SetUpConfig(enums.ProgramMode(programFlags.Mode))
	if err != nil {
		return err
	}

	//create project custom logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.NetworkConfig.Self.ProjectName)

	//starting db connection
	dbConnection := mongodb.NewMongoApi(ctx, cfg.ServiceConfig.Db.Port, cfg.ServiceConfig.Db.Ip, cfg.ServiceConfig.Db.UserName, cfg.ServiceConfig.Db.Password, cfg.ServiceConfig.Db.Name, l)
	if err != nil {
		return err
	}
	err = dbConnection.StartDbConnection()
	if err != nil {
		return err
	}
	defer dbConnection.CloseDbConnection()

	authPort := auth.NewPepperSaltAuthenticator("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		cfg.ServiceConfig.Auth.PepperLetters,
		75, cfg.ServiceConfig.Auth.PepperLength, l)

	userApplication := application.NewUserApi(dbConnection, authPort, l)

	systemCh := make(chan error)
	//start http server
	go func() {
		var s left.BaseServer = rest.NewRestServer(ctx, *cfg, l)
		err = s.ListenAndServe(userApplication)
		if err != nil {
			l.Error(err)
		}
		systemCh <- err
	}()

	//Turn on for rabbitmq connection
	// go func() {
	// 	s, err := rabbitmq.NewRabbitmqUserConsumer(l, ctx, cfg)
	// 	if err != nil {
	// 		l.Error(err)
	// 		systemCh <- err
	// 	}
	// 	err = s.ListenAndServe(userApplication)
	// 	if err != nil {
	// 		l.Error(err)
	// 	}
	// 	systemCh <- err
	// }()
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
