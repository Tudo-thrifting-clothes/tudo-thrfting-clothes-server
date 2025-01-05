package initialize

import (
	"tudo-thrfting-clothes-server/global"
	"tudo-thrfting-clothes-server/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
