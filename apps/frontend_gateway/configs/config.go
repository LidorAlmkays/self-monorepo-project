package configs

import "github.com/LidorAlmkays/self-monorepo-project/libs/configs/project_base_info"

type Config struct {
	ServiceConfig     *ServiceConfig
	BaseConfig        *project_base_info.FrontendGateway
	UserServiceConfig *project_base_info.UserService
}

type ServiceConfig struct {
	Frontend struct {
		Url string `yaml:"url" validate:"required,url" env:"FRONTEND_URL"`
	} `yaml:"frontend" validate:"required"`
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required" env:"USER_EXCHANGE_NAME"`
		Url              string `yaml:"url" validate:"required" env:"RABBITMQ_URL"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}
