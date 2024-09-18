package ports

import "github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"

type UserPort interface {
	AddUser(models.UserModel) error
}
