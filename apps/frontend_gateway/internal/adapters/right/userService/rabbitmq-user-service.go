package userService

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks"
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitmqUserService struct {
	ch  *amqp.Channel
	ctx context.Context
	l   logger.CustomLogger
	cfg configs.Config
}

func NewRabbitmqUserService(ctx context.Context, l logger.CustomLogger, cfg configs.Config) (UserServiceApi, error) {
	conn, err := frameworks.GetRabbitmqConnection(cfg.ServiceConfig.Rabbitmq.Url, l)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &rabbitmqUserService{ch, ctx, l, cfg}, nil
}

func (userService *rabbitmqUserService) AddUser(user user_dtos.AddUserDTO) error {

	ctx, cancel := context.WithTimeout(userService.ctx, 5*time.Second)
	defer cancel()
	messageBody, err := json.Marshal(user)
	if err != nil {
		userService.l.Error(errors.New("failed to marshal user data, cant send to the rabbitmq"))
		return err
	}

	err = userService.ch.PublishWithContext(ctx,
		userService.cfg.ServiceConfig.Rabbitmq.UserExchangeName, // exchange
		"user-add", // routing key
		false,      // mandatory
		false,      // immediate
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

// LoginUser implements UserServiceApi.
func (userService *rabbitmqUserService) LoginUser(user user_dtos.AuthenticateUserDTO) (*user_dtos.UserTokenResponseDTO, error) {
	panic("unimplemented LoginUser rabbitmq")
}
