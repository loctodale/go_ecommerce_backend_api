package initialize

import (
	"fmt"
	"github.com/loctodale/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Load config mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitMysql()
	InitReids()

	r := InitRouter()
	r.Run()
}
