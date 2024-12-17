package network

type AuthDB struct {
	ProjectName string `yaml:"project-name" validate:"required" env:"AUTHDB_SERVICE_PROJECT_NAME"`
	Ip          string `yaml:"ip" validate:"required" env:"AUTHDB_SERVICE_IP"`
	Port        int    `yaml:"port" validate:"required,min=1,max=65535" env:"AUTHDB_SERVICE_PORT"`
}
