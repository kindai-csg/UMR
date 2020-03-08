package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
)

type AccountController struct {
	interactor usecase.AccountInteractor
}

func NewAccountController(ldapHandler database.LdapHandler, redisHandler database.RedisHandler) *AccountController {
	accountController := AccountController {
		interactor: usecase.AccountInteractor {
			AccountRepository: &database.AccountRepository {
				LdapHandler: ldapHandler,
				RedisHandler: redisHandler,
			},
			AuthenticationCodeRepository: &database.AuthenticationCodeRepository {
				RedisHandler: redisHandler,
			},
		},
	}
	return &accountController
}

func (controller *AccountController) TemporaryCreate(c Context) {
	account := domain.Account{}
	c.Bind(&account)
	err := controller.interactor.DuplicateCheck(account.ID)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	err = controller.interactor.TemporaryRegistration(account)
	if err != nil {
		c.JSON(500, NewMsg(err.Error())) 
		return
	}
	c.JSON(200, NewMsg("仮登録が完了しました."))
} 


func (controller *AccountController) AuthenticationCreate(c Context) {
	auth := domain.AuthenticationCode{}
	c.Bind(&auth)
	err := controller.interactor.AuthenticationTemporaryAccount(auth)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	account, err := controller.interactor.FindTemporaryAccount(auth.ID)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	err = controller.interactor.Registration(account)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, NewMsg("本登録が完了しました."))
}

func (controller *AccountController) GetAllAccounts(c Context) {
	accounts, err := controller.interactor.GetAllAccounts()
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, accounts)
}
