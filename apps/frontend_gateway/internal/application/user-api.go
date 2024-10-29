package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type userApi struct {
	userService ports.UserServicePorts
	l           logger.CustomLogger
}

func NewUserApi(userManagerPorts ports.UserServicePorts, l logger.CustomLogger) ports.UserPort {
	return &userApi{
		userService: userManagerPorts,
		l:           l,
	}
}

func (uApi *userApi) AddUser(user models.UserModel) error {
	uApi.l.Info("Adding user, sending to the user service API.")
	err := uApi.userService.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}
