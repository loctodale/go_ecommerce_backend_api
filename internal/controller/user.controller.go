package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/loctodale/go-ecommerce-backend-api/internal/service"
	"github.com/loctodale/go-ecommerce-backend-api/pkg/reponse"
)

//	type UserController struct{
//		userService *service.UserService
//	}
//
//	func NewUserController() *UserController {
//		return &UserController{
//			userService: service.NewUserService(),
//		}
//	}
//
// func (uc *UserController) GetUserById(c *gin.Context) {
//
//		// reponse.SuccessReponse(c, 1, []string{"loctodale", "go learning"})
//		reponse.ErrorResponse(c, -4);
//	}
type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("random", "test")
	reponse.SuccessReponse(c, reponse.CodeSuccess, result)
}
