package rabbitmq

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqManager struct {
	cfg  configs.Config
	conn *amqp.Connection
	ctx  context.Context
	chs  []*amqp.Channel
	l    logger.CustomLogger
}

func NewRabbitmqManager(cfg configs.Config, ctx context.Context, l logger.CustomLogger) *RabbitmqManager {
	return &RabbitmqManager{cfg: cfg, ctx: ctx, l: l}
}

func (rmqM *RabbitmqManager) StartConnection() error {
	conn, err := amqp.Dial(rmqM.cfg.SharedConfig.Rabbitmq.Url)
	if err != nil {
		return err
	}
	rmqM.conn = conn
	rmqM.l.Info("Rabbitmq connection is setup.")

	return nil
}

func (rmqM *RabbitmqManager) CloseConnection() error {
	rmqM.l.Info("Closing rabbitmq connection with all chanel's")
	err := rmqM.conn.Close()
	if err != nil {
		return err
	}
	if rmqM.chs != nil {
		for _, ch := range rmqM.chs {
			if ch != nil {
				err := ch.Close()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
