package configs

import (
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs/project_base_info"
)

func SetUpConfig(debugMode bool) (*Config, error) {
	var err error = nil
	var cfg *Config = &Config{}
	if debugMode {
		cfg.BaseConfig, err = libConfigs.GetConfig[project_base_info.FrontendGateway]("../config_files/yamls/frontend-gateway.yaml", libConfigs.YAML)
		if err != nil {
			return nil, err
		}
		cfg.UserServiceConfig, err = libConfigs.GetConfig[project_base_info.UserService]("../config_files/yamls/user-service.yaml", libConfigs.YAML)
		if err != nil {
			return nil, err
		}
		cfg.ServiceConfig, err = libConfigs.GetConfig[ServiceConfig]("./configs/frontend-gateway.yaml", libConfigs.YAML)
		if err != nil {
			return nil, err
		}
	} else {

	}
	return cfg, nil
}
