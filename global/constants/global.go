package global

import (
	"gogofly/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Log *zap.SugaredLogger
	DB  *gorm.DB
  RDB *config.RedisClient
)
