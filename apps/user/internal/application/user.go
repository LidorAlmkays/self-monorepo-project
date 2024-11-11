package application

import (
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/outgoing"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application/auth"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type user struct {
	db   db.DbPort
	auth auth.AuthPort
	l    logger.CustomLogger
}

func NewUserApi(db db.DbPort, auth auth.AuthPort, l logger.CustomLogger) UserPort {
	return &user{db, auth, l}
}

func (uApi *user) AddUser(user incoming.AddUserDTO) error {
	userEntity := entities.User{
		Email:    user.Email,
		Name:     user.Name,
		BirthDay: user.BirthDay,
		UserName: user.UserName,
		Password: user.Password,
	}

	generatedPassword, err := uApi.auth.GenerateSecretPassword(&userEntity)
	if err != nil {
		uApi.l.Error(errors.New("failed to create password for user"))
		return err
	}
	userEntity.Password = generatedPassword
	err = uApi.db.AddUser(userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (uApi *user) AuthenticateUser(userData incoming.AuthenticateUserDTO) (*outgoing.UserTokenResponseDTO, error) {
	user, err := uApi.db.GetUserByEmail(userData.Email)
	if err != nil {
		return nil, err
	}
	if !uApi.auth.AuthenticateUser(userData, user) {
		err = errors.New("failed to authenticate user")
		uApi.l.Error(err)
		return nil, err
	}

	var tokenForUser string
	tokenForUser, err = uApi.auth.GenerateRandomToken()
	if err != nil {
		uApi.l.Error(errors.New("failed to generateToken, error: " + err.Error()))
		return nil, err
	}

	//TODO:(lidor) Change the role to not be only guest
	return &outgoing.UserTokenResponseDTO{
		Token: tokenForUser,
		Role:  user.Role,
	}, nil
}
