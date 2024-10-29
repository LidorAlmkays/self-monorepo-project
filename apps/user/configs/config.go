package configs

import (
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
)

type Config struct {
	SharedConfig  *libConfigs.SharedConfigs
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	Server struct {
		ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
		Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
	} `yaml:"server"  validate:"required"`
	Db struct {
		Name string `yaml:"name" validate:"required" env:"DATABASE_NAME"`
		Url  string `yaml:"url" validate:"required,url" env:"DATABASE_URL"`
	} `yaml:"db" validate:"required"`
}
