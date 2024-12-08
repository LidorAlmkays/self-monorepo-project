package project_base_info

type VideoManager struct {
	ProjectName string `yaml:"project-name"  validate:"required" env:"PROJECT_NAME"`
	Ip          string `yaml:"ip" validate:"required"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"PORT"`
} //`yaml:"video-manager" validate:"required"`
