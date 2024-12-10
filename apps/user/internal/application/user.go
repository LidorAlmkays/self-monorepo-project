package application

import (
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/adapters/right/db"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application/auth"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
)

type user struct {
	db   db.DbPort
	auth auth.AuthPort
	l    logger.CustomLogger
}

func NewUserApi(db db.DbPort, auth auth.AuthPort, l logger.CustomLogger) UserPort {
	return &user{db, auth, l}
}

func (uApi *user) AddUser(user user_dtos.AddUserDTO) error {
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

func (uApi *user) AuthenticateUser(userData user_dtos.AuthenticateUserDTO) (*user_dtos.UserTokenResponseDTO, error) {
	user, err := uApi.db.GetUserByEmail(userData.Email)
	if err != nil {
		return nil, err
	}
	if !uApi.auth.AuthenticateUser(userData, user) {
		err = errors.New("failed to authenticate user")
		uApi.l.Error(err)
		return nil, err
	}
	user.Password = userData.Password
	generatedPassword, err := uApi.auth.GenerateSecretPassword(user)
	if err != nil {
		uApi.l.Error(errors.New("failed to generate new password to user that authenticated"))
	} else {
		user.Password = generatedPassword
		err = uApi.db.UpdateUserByEmail(userData.Email, user)
		if err != nil {
			err = errors.New("failed to update user information in the database when user was authenticated")
			uApi.l.Error(err)
		}
	}
	var tokenForUser string
	tokenForUser, err = uApi.auth.GenerateRandomToken()
	if err != nil {
		uApi.l.Error(errors.New("failed to generateToken, error: " + err.Error()))
		return nil, err
	}

	return &user_dtos.UserTokenResponseDTO{
		Token: tokenForUser,
		Role:  user.Role,
	}, nil
}
