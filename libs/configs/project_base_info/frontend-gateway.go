package project_base_info

type FrontendGateway struct {
	ProjectName string `yaml:"project-name"  validate:"required" env:"FRONTEND_GATEWAY_PROJECT_NAME"`
	Ip          string `yaml:"ip" validate:"required" env:"FRONTEND_GATEWAY_IP"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"FRONTEND_GATEWAY_PORT"`
} //`yaml:"frontend-gateway" validate:"required"`
