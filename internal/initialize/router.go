package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/loctodale/go-ecommerce-backend-api/global"
	"github.com/loctodale/go-ecommerce-backend-api/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	//middlewares
	//r.Use() //logging
	//r.Use() //cross
	managerRouter := router.RouterGroupApp.Manager
	userRouter := router.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}
	return r
}
