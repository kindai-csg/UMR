package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	accountController := controllers.NewAccountController(NewLdapHandler(), NewRedisHandler())

	router.POST("/api/v1/registration/regist", func(c *gin.Context) {accountController.TemporaryCreate(c)})
	router.POST("/api/v1/registration/authentication", func(c *gin.Context) {accountController.TemporaryCreate(c)})
}