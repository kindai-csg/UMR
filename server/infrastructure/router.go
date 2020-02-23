package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/kindaidensan/UMR/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	ldapHandler := NewLdapHandler()
	redisHandler := NewRedisHandler()
	accountController := controllers.NewAccountController(ldapHandler, redisHandler)
	settingController := controllers.NewSettingController(redisHandler)


	router.POST("/admin/create_register_form", func(c *gin.Context) {settingController.CreateRegisterForm(c)})
	router.POST("/admin/get_register_form", func(c *gin.Context) {settingController.GetRegisterForm(c)})

	router.POST("/api/v1/registration/regist", func(c *gin.Context) {accountController.TemporaryCreate(c)})
	router.POST("/api/v1/registration/authentication", func(c *gin.Context) {accountController.TemporaryCreate(c)})

	Router = router
}