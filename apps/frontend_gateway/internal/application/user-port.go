package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/dtos/incoming"
)

type UserPort interface {
	RegisterUser(incoming.AddUserDTO) error
	LoginUser(incoming.AuthenticateUserDTO) (string, error)
}
