package configs

import (
	libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs/project_base_info"
)

func SetUpConfig(debugMode bool) (*Config, error) {
	var err error = nil
	var cfg *Config = &Config{}
	if debugMode {
		cfg.BaseConfig, err = libConfigs.GetConfig[project_base_info.VideoManager]("../config_files/yamls/video-manager.yaml", libConfigs.YAML)
		if err != nil {
			return nil, err
		}

	} else {

	}
	return cfg, nil
}
