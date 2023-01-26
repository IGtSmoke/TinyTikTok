package setup

import (
	"context"

	"TinyTikTok/conf"

	"github.com/go-redis/redis/v9"
)

// Rctx redis空白上下文
var Rctx = context.Background()

// Rdb redisClient
var Rdb *redis.Client

// Redis Initialize Redis client
func Redis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.RedisAddr,
		Password: conf.Conf.RedisPassword, // no password set
		DB:       conf.Conf.RedisDB,       // use default DB
	})
}
