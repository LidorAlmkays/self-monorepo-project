package frameworks

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
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
