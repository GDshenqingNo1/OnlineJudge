package global

import (
	"OnlineJudge/app/internal/model/config"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	Logger  *zap.Logger
	MysqlDB *gorm.DB
	Rdb     *redis.Client
)
