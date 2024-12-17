package auth

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/entities"
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
)

type AuthPort interface {
	GenerateSecretPassword(user *entities.User) (string, error)
	AuthenticateUser(loginInfo user_dtos.AuthenticateUserDTO, authInfo *entities.User) bool
	GenerateRandomToken() (string, error)
}
