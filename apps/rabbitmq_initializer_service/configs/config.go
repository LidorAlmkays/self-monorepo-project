package configs

import libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"

type Config struct {
	SharedConfig  *libConfigs.SharedConfigs
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	ProjectName string `yaml:"project-name"  validate:"required"`
}
