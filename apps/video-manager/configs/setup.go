package configs

import (
	"errors"

	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs/project_base_info"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/enums"
	"github.com/go-playground/validator"
)

func SetUpConfig(programMode enums.ProgramMode) (*Config, error) {
	var err error = nil
	var cfg *Config = &Config{}

	confType := libConfigs.ENV
	stringFiller := "env"
	switch programMode {
	case enums.DevelopmentMode:
		{
			stringFiller = "yaml"
			confType = libConfigs.YAML
		}
	case enums.ProductionMode:
		{
			confType = libConfigs.ENV
			stringFiller = "env"
		}
	default:
		{
			return nil, errors.New("received an invalid program mode")
		}
	}
	cfg.BaseConfig, err = libConfigs.GetConfig(&project_base_info.VideoManager{}, "../config_files/"+stringFiller+"s/video-manager."+stringFiller, confType)
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
