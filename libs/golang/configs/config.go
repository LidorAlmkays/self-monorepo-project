package configs

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func GetConfig[T any](fileFullPath string, fileType ConfigTypes) (*T, error) {
	var cfg *T
	var err error

	switch fileType {
	case YAML:
		{
			if fileFullPath != "" {
				cfg, err = openYamlConfig[T](fileFullPath)
			}
			if err != nil {
				fmt.Print("No yaml file was found.")
				return nil, err
			}
		}
	case ENV:
		{
			err = loadConfigFromEnvOrFile(cfg, fileFullPath)
			if err != nil {
				fmt.Print(err.Error())
				return nil, err
			}
		}
	default:
		{
			err := errors.New("no config file type was selected, ENV is also path variables")
			fmt.Print(err)
			return nil, err
		}
	}

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
