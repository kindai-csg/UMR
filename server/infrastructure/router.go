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
	// config取得
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// 各handler生成
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

	// 各コントローラ生成
	accountController := controllers.NewAccountController(ldapHandler, redisHandler, mailHandler, sqlHandler)
	settingController := controllers.NewSettingController(redisHandler)
	authenticationController := controllers.NewAuthenticationController(redisHandler)
	appController := controllers.NewAppController(sqlHandler)

	/*
	-------------------
	|	管理者権限API	|
	-------------------
	*/
	admin := router.Group("/admin", tokenHandler.AdminAuth)
	admin.POST("/create_register_form", func(c *gin.Context) {settingController.CreateRegisterForm(c)})
	admin.POST("/get_register_form", func(c *gin.Context) {settingController.GetRegisterForm(c)})
	admin.POST("/get_all_accounts", func(c *gin.Context) {accountController.GetAllAccounts(c)})
	admin.POST("/activation", func(c *gin.Context) {accountController.Activation(c)})
	admin.POST("/get_all_non_active_account_id", func(c *gin.Context) {accountController.GetAllNonActiveAccountID(c)})
	admin.POST("/delete_account", func(c *gin.Context) {accountController.DeleteAccount(c)})

	/*
	----------------
	|	登録系API	|
	----------------
	*/
	router.POST("/register", func(c *gin.Context) {
		err := authenticationController.AuthenticationFormToken(c)
		if err != nil {
			return
		}
		accountController.TemporaryCreate(c)
	})
	router.POST("/authentication", func(c *gin.Context) {accountController.AuthenticationCreate(c)})

	/*
	----------------
	|	認証系API	|
	----------------
	*/
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
	router.POST("/get_token_authority", func(c *gin.Context) {tokenHandler.GetTokenAuthority(c)})

	/*
	------------------
	|	ユーザー操作API	|
	------------------
	*/
	user := router.Group("/user", tokenHandler.UserAuth)
	user.POST("/create_app", func(c *gin.Context) {appController.CreateApplication(c)})
	user.POST("/get_app", func(c *gin.Context) {appController.GetApplication(c)})


	Router = router
}