package global

import (
	"github.com/loctodale/go-ecommerce-backend-api/pkg/logger"
	"github.com/loctodale/go-ecommerce-backend-api/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
