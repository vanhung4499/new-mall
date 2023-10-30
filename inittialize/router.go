package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-mall/global"
	"new-mall/middleware"
	"new-mall/router"
)

func Routers() *gin.Engine {

	// Set to publish mode
	if global.Config.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()

	InstallPlugin(Router)

	Router.StaticFS(global.Config.Local.Path, http.Dir(global.Config.Local.Path))
	//Router.Use(middleware.LoadTls())  // Open it and you can use https
	global.Log.Info("use middleware logger")
	// Cross domain origin
	Router.Use(middleware.Cors()) // If you need to cross domain, comment it
	global.Log.Info("use middleware cors")
	// Conveniently add routing group prefix for multiple servers to go online.
	//Mall admin routing
	adminRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage-api")
	PublicGroup := Router.Group("")

	{
		// health monitoring
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		// routing initialization
		adminRouter.InitManageAdminUserRouter(ManageGroup)
		adminRouter.InitManageGoodsCategoryRouter(ManageGroup)
		adminRouter.InitManageGoodsInfoRouter(ManageGroup)
		adminRouter.InitManageCarouselRouter(ManageGroup)
		adminRouter.InitManageIndexConfigRouter(ManageGroup)
		adminRouter.InitManageOrderRouter(ManageGroup)
	}
	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	{
		// 商城前端路由
		mallRouter.InitMallCarouselIndexRouter(MallGroup)
		mallRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		mallRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		mallRouter.InitMallUserRouter(MallGroup)
		mallRouter.InitMallUserAddressRouter(MallGroup)
		mallRouter.InitMallShopCartRouter(MallGroup)
		mallRouter.InitMallOrderRouter(MallGroup)
	}
	global.Log.Info("router register success")
	return Router
}
