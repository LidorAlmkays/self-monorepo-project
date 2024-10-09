package configs

import libConfigs "github.com/LidorAlmkays/self-monorepo-project/libs/golang/configs"

type Config struct {
	SharedConfig  *libConfigs.SharedConfigs
	ServiceConfig *ServiceConfig
}

type ServiceConfig struct {
	ProjectName   string   `yaml:"project-name"  validate:"required"`
	ServicesNames []string `yaml:"services-names"  validate:"required"`
	Queues        map[string]struct {
		RoutingKey string `yaml:"routing-key" validate:"required"`
		QueueName  string `yaml:"queue-name" validate:"required"`
	} `yaml:"queues" validate:"required"`
}
