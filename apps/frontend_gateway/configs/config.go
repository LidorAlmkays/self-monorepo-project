package configs

import libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"

type Config struct {
	SharedConfig  *libConfigs.SharedConfigs
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	Frontend struct {
		Url string `yaml:"url" validate:"required,url"`
	} `yaml:"frontend" validate:"required"`
}
