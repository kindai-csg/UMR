package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/controllers"
	// jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"os"
	"log"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	ldapHandler := NewLdapHandler()
	redisHandler := NewRedisHandler()
	mailHandler := NewMailHandler()
	sqlHandler := NewSqlHandler()
	if sqlHandler == nil {
		os.Exit(1);
	}

	accountController := controllers.NewAccountController(ldapHandler, redisHandler, mailHandler, sqlHandler)
	settingController := controllers.NewSettingController(redisHandler)
	authenticationController := controllers.NewAuthenticationController(redisHandler)

	authMiddleware := func(c *gin.Context) {
		log.Println("call middleware")
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("test"), nil
		});

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"Msg": "認証に失敗しました",
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
		tokenString, _ := token.SignedString([]byte("test"))
		c.JSON(200, gin.H{
			"token": tokenString,
		})
	})

	Router = router
}