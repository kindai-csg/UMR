package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/controllers"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	LdapConfig LdapConfig
	RedisConfig RedisConfig
	SqlConfig SqlConfig
	MailConfig MailConfig
	JWTConfig JWTConfig
}

type JWTConfig struct {
	Secret string  `toml:"secret"`
}

var Router *gin.Engine

func init() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	redisHandler := NewRedisHandler(config.RedisConfig)
	mailHandler := NewMailHandler(config.MailConfig)
	sqlHandler := NewSqlHandler(config.SqlConfig)
	ldapHandler := NewLdapHandler(config.LdapConfig)
	tokenHandler := NewTokenHandler(config.JWTConfig.Secret)
	if ldapHandler == nil {
		os.Exit(1)
	}
	if sqlHandler == nil {
		os.Exit(2);
	}

	accountController := controllers.NewAccountController(ldapHandler, redisHandler, mailHandler, sqlHandler)
	settingController := controllers.NewSettingController(redisHandler)
	authenticationController := controllers.NewAuthenticationController(redisHandler)

	admin := router.Group("/admin", tokenHandler.AuthMiddleware)
	admin.POST("/create_register_form", func(c *gin.Context) {settingController.CreateRegisterForm(c)})
	admin.POST("/get_register_form", func(c *gin.Context) {settingController.GetRegisterForm(c)})
	admin.POST("/get_all_accounts", func(c *gin.Context) {accountController.GetAllAccounts(c)})
	admin.POST("/activation", func(c *gin.Context) {accountController.Activation(c)})
	admin.POST("/get_all_non_active_account_id", func(c *gin.Context) {accountController.GetAllNonActiveAccountID(c)})
	admin.POST("/delete_account", func(c *gin.Context) {accountController.DeleteAccount(c)})

	router.POST("/register", func(c *gin.Context) {
		err := authenticationController.AuthenticationFormToken(c)
		if err != nil {
			return
		}
		accountController.TemporaryCreate(c)
	})

	router.POST("/authentication", func(c *gin.Context) {accountController.AuthenticationCreate(c)})


	router.POST("/login", func(c *gin.Context) {
		account := domain.LoginAccount{}
		c.Bind(&account)
		err := accountController.Login(account.ID, account.Password, account.IsAdmin)
		if err != nil {
			c.JSON(500, controllers.NewMsg(err.Error()))
			return
		}
		token := tokenHandler.CreateToken(account.ID, account.IsAdmin)
		c.JSON(200, gin.H{
			"token": token,
		})
	})

	Router = router
}