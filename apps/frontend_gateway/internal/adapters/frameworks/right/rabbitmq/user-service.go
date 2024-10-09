package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type userService struct {
	ch  *amqp.Channel
	ctx context.Context
	l   logger.CustomLogger
	cfg configs.Config
}

func (rmqM *RabbitmqManager) NewUserService() (ports.UserServicePorts, error) {
	ch, err := rmqM.conn.Channel()
	if err != nil {
		return nil, err
	}
	rmqM.chs = append(rmqM.chs, ch)
	return &userService{ch, rmqM.ctx, rmqM.l, rmqM.cfg}, nil
}

func (userService *userService) AddUser(user models.UserModel) error {

	ctx, cancel := context.WithTimeout(userService.ctx, 5*time.Second)
	defer cancel()
	messageBody, err := json.Marshal(user)
	if err != nil {
		userService.l.Error(errors.New("failed to marshal user data, cant send to the rabbitmq"))
		return err
	}
	err = userService.ch.PublishWithContext(ctx,
		userService.cfg.SharedConfig.Rabbitmq.MainExchangeName, // exchange
		"user-service.AddUser",                                 // routing key
		false,                                                  // mandatory
		false,                                                  // immediate
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
