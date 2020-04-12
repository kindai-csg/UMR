package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
	"errors"
)

type SettingController struct {
	interactor usecase.SettingInteractor
}

func NewSettingController(redisHandler database.RedisHandler) *SettingController {
	settingController := SettingController {
		interactor: usecase.SettingInteractor {
			RegisterFormRepository: &database.RegisterFormRepository {
				RedisHandler: redisHandler,
			},
		},
	}
	return &settingController
}

func (controller *SettingController) CreateRegisterForm(c Context) {
	form := domain.RegisterForm{}
	c.Bind(&form)
	form, err := controller.interactor.IssueFormToken(form.Time)
	if err != nil {
		c.JSON(500, errors.New("faild: create token"))
		return
	}
	c.JSON(200, form)
}

func (controller *SettingController) GetRegisterForm(c Context) {
	form, err := controller.interactor.GetFormToken()
	if err != nil {
		c.JSON(500, errors.New("faild: create token"))
		return
	}
	c.JSON(200, form)
}
