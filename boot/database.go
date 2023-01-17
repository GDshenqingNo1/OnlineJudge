package boot

import (
	g "OnlineJudge/app/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.Database.Mysql
	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(g.Config.Database.Mysql.GetConnMaxIdleTime())
	sqlDB.SetConnMaxIdleTime(g.Config.Database.Mysql.GetConnMaxLifeTime())
	sqlDB.SetMaxOpenConns(g.Config.Database.Mysql.MaxOpenConns)
	sqlDB.SetMaxIdleConns(g.Config.Database.Mysql.MaxIdleConns)
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatal("connect to mysql failed.", zap.Error(err))
	}
	g.MysqlDB = db
	g.Logger.Info("initialize mysql successfully!")
}

func RedisSetup() {
	config := g.Config.Database.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatal("connect to redis failed.", zap.Error(err))
	}
	g.Rdb = rdb

	g.Logger.Info("initialize redis client successfully.")

}