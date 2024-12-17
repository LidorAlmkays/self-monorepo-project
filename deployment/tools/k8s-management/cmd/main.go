package main

import (
	"flag"
	"log"
	"os"
	"strings"

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

	// ctx := context.Background()

	var err error
	// //open configs
	// cfg, err := configs.SetUpConfig(enums.ProgramMode(programFlags.Mode))
	// if err != nil {
	// 	return err
	// }

	// //create project custom logger
	// var l logger.CustomLogger = logger.NewStackedCustomLogger(cfg.NetworkConfig.Self.ProjectName)

	return err
}

func main() {
	err := setUp()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	os.Exit(0)
}
