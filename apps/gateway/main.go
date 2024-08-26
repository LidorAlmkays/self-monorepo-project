package main

import (
	"fmt"

	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

func Hello(name string, l *logger.CustomLogger) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("gateway"))
}
