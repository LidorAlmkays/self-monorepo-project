package db

import "github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"

// DbPort is the port for a db adapter
type DbPort interface {
	CloseDbConnection() error
	StartDbConnection() error
	AddUser(entities.User) error
	GetUserByEmail(string) (*entities.User, error)
	GetUserByUsernameAndPassword(string, string) (*entities.User, error)
	UpdateUserByEmail(email string, newUserData *entities.User) error
}
