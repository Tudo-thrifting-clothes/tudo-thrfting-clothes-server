package global

import (
	"tudo-thrfting-clothes-server/pkg/logger"
	"tudo-thrfting-clothes-server/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
