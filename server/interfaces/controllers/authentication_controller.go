package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
)

type AuthenticationController struct {
	interactor usecase.AuthenticationInteractor
}

func NewAuthenticationController(redisHandler database.RedisHandler) *AuthenticationController {
	authenticationController := AuthenticationController {
		interactor: usecase.AuthenticationInteractor {
			RegisterFormRepository: &database.RegisterFormRepository {
				RedisHandler: redisHandler,
			},
		},
	}
	return &authenticationController
}

func (controller *AuthenticationController) AuthenticationFormToken(c Context) error {
	form := domain.RegisterForm{}
	c.Bind(&form)
	err := controller.interactor.AuthenticationFormToken(form.Token)
	if err != nil {
		c.JSON(500, NewMsg("無効なフォームです."))
		return err;
	}
	return nil
}
