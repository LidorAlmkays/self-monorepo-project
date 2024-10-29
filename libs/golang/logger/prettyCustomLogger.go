package logger

import (
	"fmt"
	"log"
)

type prettyCustomLogger struct {
	prefix string
}

func NewPrettyCustomLogger(prefix string) CustomLogger {
	return &prettyCustomLogger{prefix: prefix}
}

// Info prints the message in yellow to the console with a timestamp
func (l *prettyCustomLogger) Info(msg string) {
	log.Printf("[INFO]  %s\n", l.prefix)
	fmt.Printf("\033[33m%s\033[0m\n", msg)
}

// // Error prints the error message in red to the console with a timestamp
func (l *prettyCustomLogger) Error(err error) {
	log.Printf("[ERROR] %s\n", l.prefix)
	panic(fmt.Sprintf("\033[31mERROR: %s\033[0m\n", err.Error()))
}

// Message prints the message in green to the console with a timestamp
func (l *prettyCustomLogger) Message(msg string) {
	log.Printf("[MESSAGE] %s\n", l.prefix)
	fmt.Printf("\033[32m%s\033[0m\n", msg)
}

// Debug prints the message in blue to the console with a timestamp
func (l *prettyCustomLogger) Debug(msg string) {
	log.Printf("[DEBUG] %s\n", l.prefix)
	fmt.Printf("\033[34m%s\033[0m\n", msg)
}
