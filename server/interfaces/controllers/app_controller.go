package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
)

type AppController struct {
	interactor usecase.AppInteractor
}

func NewAppController(sqlHandler database.SqlHandler) *AppController {
	appController := AppController {
		interactor: usecase.AppInteractor {
			AppRepository: &database.AppRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
	return &appController
}

func (controller *AppController) CreateApplication(c Context) {
	app := domain.App{}
	c.Bind(&app)
	userid, exist := c.Get("userid")
	if !exist {
		c.JSON(500, NewMsg("ユーザーIDエラー"))
		return
	}
	a, err := controller.interactor.CreateApplication(userid.(string), app)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, a)
}

func (controller *AppController) GetApplication(c Context) {
	userid, exist := c.Get("userid")
	if !exist {
		c.JSON(500, NewMsg("ユーザーIDエラー"))
		return
	}
	apps, err := controller.interactor.FindByUserId(userid.(string))
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, apps)
}