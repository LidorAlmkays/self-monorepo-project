package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/adapters/right/userService"
	userServiceIncoming "github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/enums"
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

func (uApi *userApi) RegisterUser(user incoming.AddUserDTO) error {
	uApi.l.Info("Adding user, sending to the user service API.")
	userToRegister := userServiceIncoming.AddUserDTO{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		BirthDay: user.BirthDay,
		Name:     user.Name,
		Role:     enums.User,
	}
	err := uApi.userService.AddUser(userToRegister)
	if err != nil {
		return err
	}
	return nil
}

// LoginUser implements UserPort.
func (uApi *userApi) LoginUser(user incoming.AuthenticateUserDTO) (string, error) {
	uApi.l.Info("Logging user, sending to the user service API to receive his token")
	userLogginIn := userServiceIncoming.AuthenticateUserDTO{
		Email:    user.Email,
		Password: user.Password,
	}

	UserTokenResponse, err := uApi.userService.LoginUser(userLogginIn)
	if err != nil {
		return "", err
	}
	//TODO:(Save the token in redis with its role)
	//TODO:(create logic for connecting to redis and saving it)
	return UserTokenResponse.Token, nil
}
