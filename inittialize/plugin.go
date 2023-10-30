package initialize

import (
	"fmt"
	"new-mall/global"
	"new-mall/middleware"
	"new-mall/plugin/email"
	"plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("Non Authentication plug-in installation ==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("Authentication plug-in installation==》", PrivateGroup)
	PrivateGroup.Use(middleware.UserJWTAuth()).Use(middleware.AdminJWTAuth())
	// Add a plug-in that is linked to role permissions. Example. In local sample mode and online warehouse mode, please note that the import above can be switched by yourself. The effect is the same.
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.Config.Email.To,
		global.Config.Email.From,
		global.Config.Email.Host,
		global.Config.Email.Secret,
		global.Config.Email.Nickname,
		global.Config.Email.Port,
		global.Config.Email.IsSSL,
	))
}
