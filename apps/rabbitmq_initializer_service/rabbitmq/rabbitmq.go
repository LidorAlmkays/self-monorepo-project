package rabbitmq

import (
	"context"
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/rabbitmq_initializer_service/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
)

type RabbitmqManager struct {
	cfg  configs.Config
	conn *amqp.Connection
	ctx  context.Context
	l    logger.CustomLogger
}

func NewRabbitmqManager(l logger.CustomLogger, cfg configs.Config, ctx context.Context) *RabbitmqManager {
	return &RabbitmqManager{l: l, cfg: cfg, ctx: ctx}
}

func (rmqM *RabbitmqManager) Connect() error {
	conn, err := amqp.Dial(rmqM.cfg.SharedConfig.Rabbitmq.Url)
	if err != nil {
		return err
	}
	rmqM.conn = conn
	rmqM.l.Message("Rabbitmq connection is setup.")
	return nil
}

func (rmqM RabbitmqManager) InitializeProjectRabbitMq(servicesNames []string) error {
	ch, err := rmqM.conn.Channel()
	if err != nil {
		rmqM.l.Error(errors.New("failed to set up a channel"))
		return err
	}
	rmqM.l.Message("Rabbitmq channel set up for setting up main exchange")

	err = ch.ExchangeDeclare(rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName, "topic", false, false, false, false, amqp.Table{})
	if err != nil {
		rmqM.l.Error(errors.New("failed to declare exchange with the name: " + rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName))
		return err
	}
	errs, _ := errgroup.WithContext(rmqM.ctx)
	for _, name := range servicesNames {
		errs.Go(func() error { return rmqM.createQueuesForService(name) })
		if err != nil {
			rmqM.l.Error(errors.New("failed to set up the queues for service " + name))
			return err
		}
	}

	return errs.Wait()
}

func (rmqM *RabbitmqManager) createQueuesForService(serviceName string) error {
	ch, err := rmqM.conn.Channel()
	if err != nil {
		rmqM.l.Error(errors.New("failed to set up a channel on service name: " + serviceName))
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName, //main request handler queue name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		rmqM.l.Error(errors.New("failed to declare queue with the name: " + rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName))
		return err
	}
	err = ch.QueueBind(rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName, rmqM.cfg.ServiceConfig.Queues[serviceName].RoutingKey, rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName, false, amqp.Table{})
	if err != nil {
		rmqM.l.Error(errors.New("failed to bind queue with the name: " + rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName + " ,to the exchanged named: " + rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName))
		return err
	}
	_, err = ch.QueueDeclare(
		"Response."+rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName, // Response queue name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		rmqM.l.Error(errors.New("failed to declare queue with the name: Response." + rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName))
		return err
	}
	err = ch.QueueBind("Response."+rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName, "Response."+rmqM.cfg.ServiceConfig.Queues[serviceName].RoutingKey, rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName, false, amqp.Table{})
	if err != nil {
		rmqM.l.Error(errors.New("failed to bind queue with the name: Response." + rmqM.cfg.ServiceConfig.Queues[serviceName].QueueName + " ,to the exchanged named: " + rmqM.cfg.SharedConfig.Rabbitmq.MainExchangeName))
		return err
	}
	rmqM.l.Message("The service named: " + serviceName + ", queues are ready and set up")

	return nil
}

func (rmqM *RabbitmqManager) CloseConnection() error {
	rmqM.l.Info("Closing rabbitmq connection with that is set up channel")
	err := rmqM.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
