package network

type Frontend struct {
	Ip   string `yaml:"ip" validate:"required" env:"FRONTEND_IP"`
	Port int    `yaml:"port" validate:"required,min=1,max=65535" env:"FRONTEND_PORT"`
}
