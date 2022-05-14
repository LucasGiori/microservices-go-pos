package service

import (
	"context"
	"microservices/user/internal/hash"
	"microservices/user/pkg/model"
	"time"

	"gitlab.com/pos-alfa-microservices-go/core/config"

	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"github.com/golang-jwt/jwt"
)

const (
	expirationTime = time.Hour * 1
)

type AuthManager interface {
	Login(user *model.AuthRequest) (*model.JWT, error)
}

type AuthJWT struct {
	appConfig   *config.AppConfig
	userService Service
}

func NewAuthJWT(appConfig *config.AppConfig, userService Service) AuthManager {
	return &AuthJWT{
		appConfig:   appConfig,
		userService: userService,
	}
}

func (a AuthJWT) Login(request *model.AuthRequest) (*model.JWT, error) {
	user, err := a.userService.FindByLogin(context.Background(), request.Login)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, customErrors.ErrInvalidUser
	}

	if !hash.Validate(user.Password, request.Password) {
		return nil, customErrors.ErrInvalidUser
	}

	return a.newToken(request.Login)
}

func (a AuthJWT) newToken(login string) (*model.JWT, error) {
	expiration := time.Now().Add(expirationTime)
	info := jwt.MapClaims{}
	info["authorized"] = true
	info["login"] = login
	info["exp"] = expiration.Unix()

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := jwt.SignedString([]byte(a.appConfig.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return &model.JWT{Token: token, Expiration: expiration}, nil
}
