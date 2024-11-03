package configs

type SharedConfigs struct {
	Rabbitmq struct {
		UserExchangeName string `yaml:"user-exchange-name" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
	} `yaml:"rabbitmq,omitempty" validate:"required"`
	UserService struct {
		ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
		Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
		Ip          string `yaml:"ip" validate:"required"`
	} `yaml:"user-service" validate:"required"`
	FrontendGateway struct {
		ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
		Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
		Ip          string `yaml:"ip" validate:"required"`
	} `yaml:"frontend-gateway" validate:"required"`
}
