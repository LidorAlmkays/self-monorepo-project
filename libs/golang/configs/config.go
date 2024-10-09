package configs

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator"
)

func GetConfig[T any](configFolderPath string, configFileName string) (*T, error) {
	var cfg *T
	var err error
	if configFolderPath != "" {
		cfg, err = openYamlConfig[T](configFolderPath + configFileName)
	}
	if err != nil {
		return nil, err
	}

	err = loadEnvFile(configFolderPath)
	if err != nil {
		fmt.Print("No env file was found.")
	}

	loadConfigFromEnv(cfg)

	validate := validator.New()

	err = validate.Struct(cfg)
	if err != nil {
		return nil, err
	}

	printConfigs(cfg)

	return cfg, nil
}

func printConfigs[T any](cfg *T) {
	// This is a way to print the configs
	result, _ := json.Marshal(cfg)
	fmt.Println(string(result))
}
