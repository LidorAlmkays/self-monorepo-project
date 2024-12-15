package handler

import (
	"encoding/json"
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
)

func (h *handler) AddUserConsumer(userApi application.UserPort) error {
	ch, err := h.conn.Channel()
	if err != nil {
		h.l.Error(errors.New("failed to set up a channel"))
		return err
	}
	userAddMsgs, err := ch.Consume(
		h.cfg.NetworkConfig.Self.ProjectName, // queue
		"user-add",                           // consumer
		true,                                 // auto-ack
		false,                                // exclusive
		false,                                // no-local
		false,                                // no-wait
		nil,                                  // args
	)

	go func() {
		defer ch.Close()
		for d := range userAddMsgs {
			user := user_dtos.AddUserDTO{}
			h.l.Info("Adding user: " + string(d.Body))
			err := json.Unmarshal(d.Body, &user)
			if err != nil {
				h.l.Error(errors.New("couldn't unmarshal data received from user-add queue"))
			}
			userApi.AddUser(user)
		}
	}()
	return nil
}
