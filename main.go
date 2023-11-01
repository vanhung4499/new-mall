package new_mall

import (
	"go.uber.org/zap"
	"new-mall/core"
	"new-mall/global"
	"new-mall/inittialize"
)

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.Viper = core.Viper() // Initialize Viper
	initialize.OtherInit()
	global.Log = core.Zap() // Initialize zap log
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm() // gorm connects to database
	initialize.DBList()

	core.RunWindowsServer()
}
