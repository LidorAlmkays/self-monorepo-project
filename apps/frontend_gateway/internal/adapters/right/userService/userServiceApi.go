package userService

import "github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"

type UserServiceApi interface {
	AddUser(user models.UserRegisterModel) error
}
