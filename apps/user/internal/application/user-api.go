package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/ports"
)

type userApi struct {
	db ports.DbPort
}

func NewUserApi(db ports.DbPort) ports.UserPort {
	return &userApi{db: db}
}

func (uApi *userApi) AddUser(user models.UserModel) error {
	err := uApi.db.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}
