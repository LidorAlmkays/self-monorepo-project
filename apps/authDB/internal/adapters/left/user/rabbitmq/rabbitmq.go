package rabbitmq

import (
	"context"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/adapters/frameworks"
	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/adapters/left"
	handler "github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/adapters/left/user/rabbitmq/handlers"
	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
)

type rabbitmq struct {
	conn    *amqp.Connection
	l       logger.CustomLogger
	userApi application.UserPort
	ctx     context.Context
	cfg     configs.Config
}

func NewRabbitmqUserConsumer(l logger.CustomLogger,
	ctx context.Context,
	cfg configs.Config) (left.BaseServer, error) {
	conn, err := frameworks.GetRabbitmqConnection(cfg.ServiceConfig.Rabbitmq.Url, l)
	if err != nil {
		return nil, err
	}
	return &rabbitmq{conn: conn, l: l, ctx: ctx, cfg: cfg}, nil
}

func (r *rabbitmq) ListenAndServe(userApi application.UserPort) error {
	ch, err := r.conn.Channel()
	if err != nil {
		r.l.Error(errors.New("failed to set up a channel"))
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(r.cfg.ServiceConfig.Rabbitmq.UserExchangeName, "topic", false, false, false, false, amqp.Table{})
	if err != nil {
		r.l.Error(errors.New("failed to declare exchange with the name: " + r.cfg.ServiceConfig.Rabbitmq.UserExchangeName))
		return err
	}
	_, err = ch.QueueDeclare(
		r.cfg.NetworkConfig.Self.ProjectName, //main request handler queue name
		true,                                 // durable
		false,                                // delete when unused
		false,                                // exclusive
		false,                                // no-wait
		nil,                                  // arguments
	)
	if err != nil {
		r.l.Error(errors.New("failed to declare queue with the name: " + r.cfg.NetworkConfig.Self.ProjectName))
		return err
	}
	routingKeys := []string{"user-add", "user-update", "user-get", "user-delete"}
	for _, routingKey := range routingKeys {
		err = ch.QueueBind(r.cfg.NetworkConfig.Self.ProjectName, routingKey, r.cfg.ServiceConfig.Rabbitmq.UserExchangeName, false, amqp.Table{})
		if err != nil {
			r.l.Error(errors.New("failed to bind queue with the name: " + r.cfg.NetworkConfig.Self.ProjectName + " ,to the exchanged named: " + r.cfg.ServiceConfig.Rabbitmq.UserExchangeName))
			return err
		}
	}
	r.l.Message("set up queue for user request from rabbitmq")
	var forever chan struct{}
	h := handler.NewHandler(r.conn, r.l, userApi, r.ctx, r.cfg)
	h.AddUserConsumer(userApi)
	<-forever
	return nil
}
