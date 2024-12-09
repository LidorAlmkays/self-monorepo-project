package project_base_info

type UserService struct {
	ProjectName string `yaml:"project-name"  validate:"required" env:"USER_SERVICE_PROJECT_NAME"`
	Ip          string `yaml:"ip" validate:"required" env:"USER_SERVICE_IP"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"USER_SERVICE_PORT"`
} //`yaml:"user-service" validate:"required"`
