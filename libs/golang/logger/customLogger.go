package logger

// CustomLogger defines the interface for our custom logger
type CustomLogger interface {
	Info(msg string)
	Message(msg string)
	Debug(msg string)
	Error(err error)
}
