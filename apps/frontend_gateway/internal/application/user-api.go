package application

import (
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/ports"
)

type userApi struct {
}

func NewUserApi() ports.UserPort {
	return &userApi{}
}

func (uApi *userApi) AddUser(user models.UserModel) error {
	err := errors.New("No implemented version")
	if err != nil {
		return err
	}
	return nil
}
