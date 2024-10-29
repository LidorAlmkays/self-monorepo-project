package ports

import "github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"

type UserPort interface {
	AddUser(models.UserModel) error
}
