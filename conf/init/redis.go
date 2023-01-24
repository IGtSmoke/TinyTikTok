package init

import (
	"TinyTikTok/conf"
	"context"
	"github.com/go-redis/redis/v9"
)

// Rctx redis空白上下文
var Rctx = context.Background()

// Rdb redisClient
var Rdb *redis.Client

func Redis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword, // no password set
		DB:       conf.RedisDB,       // use default DB
	})

}
