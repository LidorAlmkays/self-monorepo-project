package configs

import (
	"os"

	"github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/adapters/frameworks/right/db"
	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port" validate:"required,min=1,max=65535"`
	} `yaml:"server"  validate:"required"`
	Db struct {
		Port     int       `yaml:"port" validate:"required,min=1,max=65535"`
		Ip       string    `yaml:"ip" validate:"required,ipv4"`
		Type     db.DbType `yaml:"type" validate:"required,oneof=mongo"`
		DbName   string    `yaml:"db_name" validate:"required"`
		Password string    `yaml:"password" validate:"required"`
		Username string    `yaml:"username" validate:"required"`
	} `yaml:"db" validate:"required"`
}

func OpenYamlConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return nil, err
	}

	//This is a way to print the configs
	// result, _ := json.Marshal(cfg)
	// fmt.Println(string(result))
	return &cfg, nil
}

// func OpenYamlConfig(path string) (*Config, error) {
// 	b, err :=  os.ReadFile(path)
// 	if err != nil {
// 		return nil,err
// 	}

// 	var cfg Config
// 	err = yaml.UnmarshalStrict(b,&cfg)
// 	if err != nil {
// 		return nil,err
// 	}

// 	return &cfg, nil
// }
