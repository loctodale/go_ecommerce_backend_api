package router

import (
	"github.com/loctodale/go-ecommerce-backend-api/internal/router/manager"
	"github.com/loctodale/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
