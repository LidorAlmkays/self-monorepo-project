package userService

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/frameworks"
	userIncoming "github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/outgoing"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitmqUserService struct {
	ch  *amqp.Channel
	ctx context.Context
	l   logger.CustomLogger
	cfg configs.Config
}

func NewRabbitmqUserService(ctx context.Context, l logger.CustomLogger, cfg configs.Config) (UserServiceApi, error) {
	conn, err := frameworks.GetRabbitmqConnection(cfg.SharedConfig.Rabbitmq.Url, l)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &rabbitmqUserService{ch, ctx, l, cfg}, nil
}

func (userService *rabbitmqUserService) AddUser(user userIncoming.AddUserDTO) error {

	ctx, cancel := context.WithTimeout(userService.ctx, 5*time.Second)
	defer cancel()
	messageBody, err := json.Marshal(user)
	if err != nil {
		userService.l.Error(errors.New("failed to marshal user data, cant send to the rabbitmq"))
		return err
	}

	err = userService.ch.PublishWithContext(ctx,
		userService.cfg.SharedConfig.Rabbitmq.UserExchangeName, // exchange
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
func (userService *rabbitmqUserService) LoginUser(user userIncoming.AuthenticateUserDTO) (*outgoing.UserTokenResponseDTO, error) {
	panic("unimplemented LoginUser rabbitmq")
}
