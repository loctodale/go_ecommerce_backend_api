package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/loctodale/go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandle()

	userRouterPublic := Router.Group("/admin")
	{
		userRouterPublic.POST("/register", userController.Register)
	}
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.GET("/active_user")
	}

}
