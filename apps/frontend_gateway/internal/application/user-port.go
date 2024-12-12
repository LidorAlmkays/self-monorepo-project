package application

import "github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"

type UserPort interface {
	RegisterUser(user_dtos.RequestToAddUserDTO) error
	LoginUser(user_dtos.AuthenticateUserDTO) (string, error)
}
