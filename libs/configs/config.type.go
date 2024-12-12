package configs

type ConfigTypes string

const (
	YAML ConfigTypes = "yaml"
	ENV  ConfigTypes = "env"
)

func (c ConfigTypes) String() string {
	return string(c)
}
