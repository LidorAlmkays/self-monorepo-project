package configs

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/configs/network"
)

type Config struct {
	NetworkConfig *NetworkCOnfig
}

type NetworkCOnfig struct {
	Self *network.VideoManager `yaml:"video-manager" validate:"required"`
}
