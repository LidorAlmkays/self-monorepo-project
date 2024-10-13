package configs

type SharedConfigs struct {
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
}
