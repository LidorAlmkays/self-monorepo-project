package configs

import (
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
)

type Config struct {
	SharedConfig  *libConfigs.SharedConfigs
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	Auth struct {
		PepperLetters string `yaml:"pepper-letters" validate:"required"`
		PepperLength  int    `yaml:"pepper-length" validate:"required"`
	} `yaml:"auth" validate:"required"`
	Db struct {
		Name string `yaml:"name" validate:"required"`
		Url  string `yaml:"url" validate:"required,url"`
	} `yaml:"db" validate:"required"`
}
