package userService

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/outgoing"
)

type UserServiceApi interface {
	AddUser(user incoming.AddUserDTO) error
	LoginUser(user incoming.AuthenticateUserDTO) (*outgoing.UserTokenResponseDTO, error)
}
