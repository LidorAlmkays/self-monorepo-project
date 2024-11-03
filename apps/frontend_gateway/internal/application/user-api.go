package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/right/userService"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type userApi struct {
	userService userService.UserServiceApi
	l           logger.CustomLogger
}

func NewUserApi(userManagerPorts userService.UserServiceApi, l logger.CustomLogger) UserPort {
	return &userApi{
		userService: userManagerPorts,
		l:           l,
	}
}

func (uApi *userApi) RegisterUser(user models.UserRegisterModel) error {
	uApi.l.Info("Adding user, sending to the user service API.")
	err := uApi.userService.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

// LoginUser implements UserPort.
func (uApi *userApi) LoginUser(user models.UserLoginModel) (string, error) {
	uApi.l.Info("Logging user, sending to the user service API to recive his token")
	return "", nil
}
