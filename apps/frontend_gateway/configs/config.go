package configs

type Config struct {
	Frontend struct {
		Url string `yaml:"url" validate:"required,url" env:"FRONTEND_URL"`
	} `yaml:"frontend" validate:"required"`
	Server struct {
		ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
		Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
	} `yaml:"server"  validate:"required"`
	Rabbitmq struct {
		Url string `yaml:"url" validate:"required" env:"RABBIT_MQ_URL"`
	} `yaml:"rabbitmq" validate:"required"`
}
