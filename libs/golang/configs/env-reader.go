package configs

import (
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func loadConfigFromEnv(cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Handle nested structs
		if field.Type.Kind() == reflect.Struct {
			loadConfigFromEnv(value.Addr().Interface())
		}

		// Process the `env` tag
		envTag := field.Tag.Get("env")
		if envTag == "" {
			continue
		}

		envValue, exists := os.LookupEnv(envTag)
		if !exists {
			continue
		}

		// Set the field value based on its type
		switch field.Type.Kind() {
		case reflect.String:
			value.SetString(envValue)
		case reflect.Int:
			intValue, err := strconv.Atoi(envValue)
			if err != nil {
				return
			}
			value.SetInt(int64(intValue))
		// Handle other types as needed
		default:
			return
		}
	}

}

func loadEnvFile(configFolderPath string) error {
	err := godotenv.Load(configFolderPath + ".env")
	if err != nil {
		return err
	}
	return nil
}
