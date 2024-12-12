package userService

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
)

type UserServiceApi interface {
	AddUser(user user_dtos.AddUserDTO) error
	LoginUser(user user_dtos.AuthenticateUserDTO) (*user_dtos.UserTokenResponseDTO, error)
}
