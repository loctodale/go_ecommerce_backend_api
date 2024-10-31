package initialize

import (
	"github.com/loctodale/go-ecommerce-backend-api/global"
	"github.com/loctodale/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger()

}
