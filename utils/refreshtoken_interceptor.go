package utils

import (
	"TinyTikTok/conf/setup"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

// Login 用户登录Token信息
type Login struct {
	Token string `form:"token" json:"token" uri:"token" xml:"token"`
}

// RefreshTokenInterceptor 刷新token(有token刷新,无token直接放过)
func RefreshTokenInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login Login

		if err := ctx.ShouldBind(&login); err != nil {
			log.Err(err)
		}

		// 不存在token
		if login.Token == "" {
			ctx.Next()
			return
		}

		tokenKey := LoginUserKey + login.Token
		// 取出userId
		tmp, err := setup.Rdb.HGet(setup.Rctx, tokenKey, "userId").Result()
		if err != nil {
			log.Err(err)
			ctx.Next()
			return
		}

		if tmp == "" {
			ctx.Next()
			return
		}
		// 将userId存入context
		userId, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			log.Err(err)
		}
		ctx.Set("userId", userId)
		// 刷新token有效期
		setup.Rdb.Expire(setup.Rctx, tokenKey, LoginUserTTL)
		ctx.Next()
	}
}
