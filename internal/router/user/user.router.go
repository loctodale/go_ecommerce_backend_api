package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limit)
	//userRouterPrivate.Use(Authen)
	//userRouterPrivate.Use(Permission)
	{
		userRouterPrivate.GET("/get_info/:id")
	}
}
