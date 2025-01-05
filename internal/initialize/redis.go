package initialize

import (
	"context"
	"fmt"
	"tudo-thrfting-clothes-server/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: r.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis connect error: ", zap.Error(err))
	}

	fmt.Printf("Redis connected on %s:%d\n", r.Host, r.Port)
	global.Rdb = rdb

	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", "100", 0).Err()

	if err != nil {
		global.Logger.Error("Redis set error: ", zap.Error(err))
	}

	val, err := global.Rdb.Get(ctx, "score").Result()

	if err != nil { // redis.Nil
		global.Logger.Error("Redis get error: ", zap.Error(err))
	}

	global.Logger.Info("key", zap.String("key", val))
}
