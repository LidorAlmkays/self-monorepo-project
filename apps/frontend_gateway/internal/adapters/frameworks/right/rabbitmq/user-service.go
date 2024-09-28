package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type userService struct {
	ch  *amqp.Channel
	ctx context.Context
	l   logger.CustomLogger
	q   *amqp.Queue
}

func (rmqM *RabbitmqManager) NewUserService() (ports.UserServicePorts, error) {
	ch, err := rmqM.Conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		"user-service", // name
		false,          // durable
		true,           // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		rmqM.l.Error(errors.New("failed to declare user-service queue"))
		return nil, err
	}

	return &userService{ch, rmqM.ctx, rmqM.l, &q}, nil
}

func (userService *userService) AddUser(user models.UserModel) error {

	ctx, cancel := context.WithTimeout(userService.ctx, 5*time.Second)
	defer cancel()
	messageBody, err := json.Marshal(user)
	if err != nil {
		userService.l.Error(errors.New("failed to parse the data into a json to send throw the rabbitmq"))
		return err
	}
	err = userService.ch.PublishWithContext(ctx,
		"",                 // exchange
		userService.q.Name, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messageBody,
		})
	if err != nil {
		userService.l.Error(errors.New("failed to publish the data to rabbitmq"))
		return err
	}

	return nil
}
