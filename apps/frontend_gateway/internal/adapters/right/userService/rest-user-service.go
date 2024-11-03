package userService

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/configs"
	"github.com/LidorAlmkays/self-monorepo-project/apps/frontend_gateway/internal/models"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type restUserService struct {
	ctx            context.Context
	l              logger.CustomLogger
	cfg            configs.Config
	userServiceUrl string
}

func NewRestUserService(ctx context.Context, l logger.CustomLogger, cfg configs.Config) (UserServiceApi, error) {
	userServiceUrl := "http://" + cfg.SharedConfig.UserService.Ip + ":" + strconv.Itoa(cfg.SharedConfig.UserService.Port)
	return &restUserService{ctx, l, cfg, userServiceUrl}, nil
}

func (r *restUserService) AddUser(user models.UserRegisterModel) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	buf := bytes.NewReader(jsonData)
	_, err = http.Post(r.userServiceUrl+"/user", "application/json", buf)
	if err != nil {
		return err
	}
	return nil
}
