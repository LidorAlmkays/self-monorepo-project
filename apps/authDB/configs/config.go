package configs

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/configs/network"
)

type Config struct {
	ServiceConfig *ServiceConfig
	NetworkConfig *NetworkConfig
}

type ServiceConfig struct {
	Auth struct {
		PepperLetters string `yaml:"pepper-letters" validate:"required" env:"PEPPER_LETTERS"`
		PepperLength  int    `yaml:"pepper-length" validate:"required" env:"PEPPER_LENGTH"`
	} `yaml:"auth" validate:"required"`
	Db struct {
		Name     string `yaml:"name" validate:"required" env:"DB_NAME"`
		UserName string `yaml:"username" validate:"required" env:"DB_USERNAME"`
		Password string `yaml:"password" validate:"required" env:"DB_PASSWORD"`
		Port     int    `yaml:"port" validate:"required,min=1,max=65535" env:"DB_PORT"`
		Ip       string `yaml:"ip" validate:"required" env:"DB_IP"`
	} `yaml:"database" validate:"required"`
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required" env:"USER_EXCHANGE_NAME"`
		Url              string `yaml:"url" validate:"required" env:"RABBITMQ_URL"`
	} `yaml:"rabbitmq" validate:"required"`
}

type NetworkConfig struct {
	Self *network.AuthDB `yaml:"authDB" validate:"required"`
}
