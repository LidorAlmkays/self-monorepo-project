package auth

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
)

type AuthPort interface {
	GenerateSecretPassword(user *entities.User) (string, error)
	AuthenticateUser(loginInfo incoming.AuthenticateUserDTO, authInfo *entities.User) bool
	GenerateRandomToken() (string, error)
}
