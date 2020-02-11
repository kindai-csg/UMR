package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
	"errors"
)

type AccountController struct {
	interactor usecase.AccountInteractor
}

func NewAccountController(ldapHandler database.LdapHandler, redisHandler RedisHandler) *AccountController {
	accountController := AccountController {
		interactor: &usecase.AccountInteractor {
			accountRepository: &database.AccountRepository {
				ldapHandler: ldapHandler,
				redisHandler: redisHandler,
			},
			authenticationCodeRepository: &database.AuthenticationCodeRepository {
				redisHandler: redisHandler,
			},
		},
	}
	return &accountController
}

func (controller *AccountController) TemporaryCreate(c Context) {
	account := domain.Account{}
	c.Bind(&account)
	err := controller.interactor.TemporaryRegistration(account)
	if err != nil {
		c.JSON(500, errors.New("faild: create"))
		return
	}
	c.JSON(200)
} 


func (controller *AccountController) AuthenticationCreate(c Context) {
	auth := domain.AuthenticationCode{}
	c.Bind(&auth)
	err := controller.interactor.AuthenticationTemporaryAccount(auth)
	if err != nil {
		c.JSON(500, errors.New("faild: certification"))
		return
	}
	account, err := controller.interactor.FindTemporaryAccount(auth.ID)
	if err != nil {
		c.JSON(500, errors.New("faild: get account information"))
		return
	}
	err := controller.interactor.Registration(account)
	if err != nil {
		c.JSON(500, errors.New("faild: create account"))
		return
	}
	c.JSON(200)
}
