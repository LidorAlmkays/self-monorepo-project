package handler

import (
	"context"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type handler struct {
	conn    *amqp.Connection
	l       logger.CustomLogger
	userApi application.UserPort
	ctx     context.Context
	cfg     configs.Config
}

func NewHandler(conn *amqp.Connection,
	l logger.CustomLogger,
	userApi application.UserPort,
	ctx context.Context,
	cfg configs.Config) *handler {

	return &handler{l: l, cfg: cfg, conn: conn, userApi: userApi, ctx: ctx}
}
