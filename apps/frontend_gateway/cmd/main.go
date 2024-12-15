package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/left/rest"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/right/userService"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/enums"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
	"github.com/LidorAlmkays/self-monorepo-project/libs/validators"
	"github.com/go-playground/validator"
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

	//create project api with the gui
	var userServiceApi userService.UserServiceApi
	userServiceApi, err = userService.NewRestUserService(ctx, l, *cfg)
	if err != nil {
		return err
	}
	userApplication := application.NewUserApi(userServiceApi, l)

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
