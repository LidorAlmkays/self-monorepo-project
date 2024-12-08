package project_base_info

type UserService struct {
	ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
	Ip          string `yaml:"ip" validate:"required"`
} //`yaml:"user-service" validate:"required"`
