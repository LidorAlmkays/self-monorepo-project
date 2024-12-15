package configs

import (
	"errors"

	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/enums"
	"github.com/go-playground/validator"
)

func SetUpConfig(programMode enums.ProgramMode) (*Config, error) {
	var err error = nil
	var cfg *Config = &Config{}

	confType := libConfigs.ENV
	basePath := ""
	switch programMode {
	case enums.DevelopmentMode:
		{
			basePath = "../configs/"
			confType = libConfigs.YAML
		}
	case enums.ProductionMode:
		{
			basePath = "../../deployment/configs/"
			confType = libConfigs.ENV
		}
	default:
		{
			return nil, errors.New("received an invalid program mode")
		}
	}
	cfg, err = libConfigs.GetConfig(cfg, basePath+"network-info."+confType.String(), confType)
	if err != nil {
		return nil, err
	}
	validate := validator.New()

	err = validate.Struct(cfg)
	if err != nil {
		err = errors.New("failed to validate config object, error: " + err.Error())
		return nil, err
	}

	return cfg, nil
}
