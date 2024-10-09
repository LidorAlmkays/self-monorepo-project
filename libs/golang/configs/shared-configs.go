package configs

type SharedConfigs struct {
	Rabbitmq struct {
		MainExchangeName string `yaml:"main-exchange-name" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}
