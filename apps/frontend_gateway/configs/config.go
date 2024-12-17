package configs

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/configs/network"
)

type Config struct {
	ServiceConfig *ServiceConfig
	NetworkConfig *NetworkConfig
}

type ServiceConfig struct {
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required" env:"USER_EXCHANGE_NAME"`
		Url              string `yaml:"url" validate:"required" env:"RABBITMQ_URL"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}

type NetworkConfig struct {
	Self     *network.FrontendGateway `yaml:"frontend-gateway" validate:"required"`
	AuthDB   *network.AuthDB          `yaml:"authDB-service" validate:"required"`
	Frontend *network.Frontend        `yaml:"frontend" validate:"required"`
}
