package ports

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"
)

// DbPort is the port for a db adapter
type DbPort interface {
	CloseDbConnection() error
	StartDbConnection() error
	AddUser(models.UserModel) error
}
