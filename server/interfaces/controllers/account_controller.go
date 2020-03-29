package controllers

import (
	"github.com/kindaidensan/UMR/usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
)

type AccountController struct {
	interactor usecase.AccountInteractor
	mail MailHandler
}

func NewAccountController(ldapHandler database.LdapHandler, redisHandler database.RedisHandler, mailHandler MailHandler) *AccountController {
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
		mail: mailHandler,
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
	code, err := controller.interactor.TemporaryRegistration(account)
	if err != nil {
		c.JSON(500, NewMsg(err.Error())) 
		return
	}
	authUrl := "https://localhost:3080/authentication?id="+account.ID+"&code="+code
	subject := "[近畿大学電子計算機研究会]メール認証"
	body := "リンク先にアクセスして認証を完了させてください\r\n"+authUrl
	err = controller.mail.SendMail(account.EmailAddress, subject, body)
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
	// account, err := controller.interactor.FindTemporaryAccount(auth.ID)
	// if err != nil {
	// 	c.JSON(500, NewMsg(err.Error()))
	// 	return
	// }
	// err = controller.interactor.Registration(account)
	// if err != nil {
	// 	c.JSON(500, NewMsg(err.Error()))
	// 	return
	// }
	// c.JSON(200, NewMsg("本登録が完了しました."))
	c.JSON(200, NewMsg("認証が完了しました."))
}

func (controller *AccountController) Activation(c Context) {
	id := c.PostForm("ID")
	err := controller.interactor.AuthenticationCheck(id)
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	account, err := controller.interactor.FindTemporaryAccount(id)
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

func (controller *AccountController) GetAllNonActiveAccountID(c Context) {
	accounts, err := controller.interactor.GetAllNonActiveAccountID()	
	if err != nil {
		c.JSON(500, NewMsg(err.Error()))
		return
	}
	c.JSON(200, accounts)
}