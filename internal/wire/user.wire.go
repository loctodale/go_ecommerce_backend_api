//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/loctodale/go-ecommerce-backend-api/internal/controller"
	"github.com/loctodale/go-ecommerce-backend-api/internal/repo"
	"github.com/loctodale/go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandle() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
