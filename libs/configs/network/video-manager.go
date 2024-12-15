package network

type VideoManager struct {
	ProjectName string `yaml:"project-name" validate:"required" env:"VIDEO_MANAGER_PROJECT_NAME"`
	Ip          string `yaml:"ip" validate:"required" env:"VIDEO_MANAGER_IP"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"VIDEO_MANAGER_PORT"`
} //`yaml:"video-manager" validate:"required"`
