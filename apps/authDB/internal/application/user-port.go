package application

import "github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"

type UserPort interface {
	AddUser(user_dtos.AddUserDTO) error
	AuthenticateUser(user_dtos.AuthenticateUserDTO) (*user_dtos.UserTokenResponseDTO, error)
}
