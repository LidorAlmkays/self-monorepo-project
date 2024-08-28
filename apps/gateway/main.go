package main

import (
	"libs/golang/logger"
)

func main() {
	var l logger.CustomLogger = logger.NewStackedCustomLogger("gateway")
	l.Message("Starting gateway")
}
