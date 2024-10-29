package application

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"
)

type userApi struct {
	db db.DbPort
	// userRsponseManager userResponseManager.UserResponseManager
}

func NewUserApi(db db.DbPort) UserPort { //, userRsponseManager userResponseManager.UserResponseManager
	return &userApi{db: db} //, userRsponseManager: userRsponseManager}
}

func (uApi *userApi) AddUser(user models.UserModel) error {
	err := uApi.db.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}
