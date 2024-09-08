package main

import (
	"libs/golang/logger"
)

//acts like an init function
func setUp(l logger.CustomLogger){
	//connect to redis database
	
}

func main() {
	//create project logger
	var l logger.CustomLogger = logger.NewStackedCustomLogger("gateway")
	l.Message("Starting gateway")
	setUp(l)
}
