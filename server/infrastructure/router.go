package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/BurntSushi/toml"
	"strings"
	"time"
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
	if ldapHandler == nil {
		os.Exit(1)
	}
	if sqlHandler == nil {
		os.Exit(2);
	}

	accountController := controllers.NewAccountController(ldapHandler, redisHandler, mailHandler, sqlHandler)
	settingController := controllers.NewSettingController(redisHandler)
	authenticationController := controllers.NewAuthenticationController(redisHandler)

	authMiddleware := func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		
		t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTConfig.Secret), nil
		})
		
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"Msg": "認証に失敗しました",
			})
			return
		}

		claims := t.Claims.(jwt.MapClaims)
		now := time.Now().Add(time.Hour * 0).Unix()
		if (claims["exp"].(float64) < float64(now)) {
			c.AbortWithStatusJSON(500, gin.H{
				"Msg": "有効期限切れです",
			})
			return
		}
		
	}

	admin := router.Group("/admin", authMiddleware)
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
		account := domain.AdminAccount{}
		c.Bind(&account)
		err := accountController.Login(account.ID, account.Password)
		if err != nil {
			c.JSON(500, controllers.NewMsg(err.Error()))
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["ID"] = account.ID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString([]byte(config.JWTConfig.Secret))
		c.JSON(200, gin.H{
			"token": tokenString,
		})
	})

	Router = router
}