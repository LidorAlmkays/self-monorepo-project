package configs

import "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs/project_base_info"

type Config struct {
	ServiceConfig *ServiceConfig
	BaseConfig    *project_base_info.UserService
}

type ServiceConfig struct {
	Auth struct {
		PepperLetters string `yaml:"pepper-letters" validate:"required"`
		PepperLength  int    `yaml:"pepper-length" validate:"required"`
	} `yaml:"auth" validate:"required"`
	Db struct {
		Name     string `yaml:"name" validate:"required"`
		UserName string `yaml:"username" validate:"required"`
		Password string `yaml:"password" validate:"required"`
		Port     int    `yaml:"port" validate:"required,min=1,max=65535"`
		Ip       string `yaml:"ip" validate:"required"`
	} `yaml:"db" validate:"required"`
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}
