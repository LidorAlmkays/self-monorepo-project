package ports

import "github.com/LidorAlmkays/self-monorepo-project/apps/gateway/internal/models"

// DbPort is the port for a db adapter
type DbPort interface {
	CloseDbConnection()
	AddUser(models.UserModel) error
}
