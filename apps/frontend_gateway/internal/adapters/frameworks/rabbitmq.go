package frameworks

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

var conn *amqp.Connection

func GetRabbitmqConnection(url string, l logger.CustomLogger) (*amqp.Connection, error) {
	var err error = nil
	var once sync.Once
	once.Do(func() {
		conn, err = amqp.Dial(url)
		l.Message("Rabbitmq connection is setup.")
	})
	return conn, err
}

func CloseRabbitmqConnection(l logger.CustomLogger) error {
	l.Info("Closing rabbitmq connection with all chanel's")
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
