package application

import "github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"

type UserPort interface {
	RegisterUser(models.UserRegisterModel) error
	LoginUser(models.UserLoginModel) (string, error)
}
