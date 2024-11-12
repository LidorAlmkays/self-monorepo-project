package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/outgoing"
)

type UserPort interface {
	AddUser(incoming.AddUserDTO) error
	AuthenticateUser(incoming.AuthenticateUserDTO) (*outgoing.UserTokenResponseDTO, error)
}
