package configs

import "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs/project_base_info"

type Config struct {
	ServiceConfig     *ServiceConfig
	BaseConfig        *project_base_info.FrontendGateway
	UserServiceConfig *project_base_info.UserService
}

type ServiceConfig struct {
	Frontend struct {
		Url string `yaml:"url" validate:"required,url"`
	} `yaml:"frontend" validate:"required"`
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}
