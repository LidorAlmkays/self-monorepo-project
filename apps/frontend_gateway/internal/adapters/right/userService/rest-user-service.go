package userService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/libs/dtos/user_dtos"
	"github.com/LidorAlmkays/self-monorepo-project/libs/logger"
)

type restUserService struct {
	ctx            context.Context
	l              logger.CustomLogger
	cfg            configs.Config
	userServiceUrl string
}

func NewRestUserService(ctx context.Context, l logger.CustomLogger, cfg configs.Config) (UserServiceApi, error) {
	userServiceUrl := "http://" + cfg.UserServiceConfig.Ip + ":" + strconv.Itoa(cfg.UserServiceConfig.Port)
	return &restUserService{ctx, l, cfg, userServiceUrl}, nil
}

func (r *restUserService) AddUser(user user_dtos.AddUserDTO) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	buf := bytes.NewReader(jsonData)
	resp, err := http.Post(r.userServiceUrl+"/user/add", "application/json", buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to register user into database")
	}
	return nil
}

func (r *restUserService) LoginUser(user user_dtos.AuthenticateUserDTO) (*user_dtos.UserTokenResponseDTO, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(jsonData)

	resp, err := http.Post(r.userServiceUrl+"/user/authenticate", "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		r.l.Info("User trying to authenticate failed")
		return nil, errors.New("failed to authenticate user")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	var tokenResponse user_dtos.UserTokenResponseDTO
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		fmt.Println("Error parsing response JSON:", err)
		return nil, err
	}

	return &tokenResponse, nil
}
