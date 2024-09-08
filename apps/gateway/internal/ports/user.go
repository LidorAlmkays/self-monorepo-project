package ports

import domain "github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/models"

type UserPort interface {
	AddUser(domain.UserModel)error
}