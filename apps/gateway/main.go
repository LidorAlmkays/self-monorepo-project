package main

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

func main() {
	var l logger.CustomLogger = logger.NewStackedCustomLogger("gateway")
	l.Message("Starting gateway")
}
